package service

import (
    "os"
    "strings"
    "strconv"
    "time"
    "github.com/google/uuid"
    "github.com/jmnoland/identity/request"
    "github.com/jmnoland/identity/model"
)

var expireHours = os.Getenv("IDENTITY_SESSION_EXPIRES");

func findUserPermissions(apps []model.Application, appName string) ([]string) {
    for i := range apps {
        if apps[i].Name == appName {
            return apps[i].Permissions
        }
    }
    return []string {}
}

func AuthenticateSession(req request.AuthenticateRequest) (model.ServiceResponse) {
    cred := GetCredential(req.Identifier)

    parts := strings.Split(cred.Secret, ":")
    outputHash := CreateSecretHashFromExisting(req.Secret, parts[1])

    if (outputHash == parts[0]) {
        existingSession := GetSessionForUserId(cred.UserId)
        if (time.Now().Before(existingSession.ExpiresAt)) {
            return CreateResponse("OK", existingSession)
        }
        usr := GetUser(cred.ClientId, cred.UserId)

        hoursToAdd, _ := strconv.ParseInt(expireHours, 10, 32)
        expiresAt := time.Now().Add(time.Hour * time.Duration(hoursToAdd))

        sess := model.Session{
            ID:             uuid.New(),
            UserId:         usr.ID,
            Permissions:    findUserPermissions(usr.Applications, req.Application),
            ExpiresAt:      expiresAt,
        }
        AddSessionCache(sess)

        return CreateResponse("OK", sess)
    }
    return CreateResponse("BADREQUEST", req)
}

func AuthenticateApiKey(req request.AuthenticateRequest) (model.ServiceResponse) {
    cred := GetCredential(req.Identifier)

    parts := strings.Split(cred.Secret, ":")
    outputHash := CreateSecretHashFromExisting(req.Secret, parts[1])

    if (outputHash == parts[0]) {
        usr := GetUser(cred.ClientId, cred.UserId)

        return CreateResponse("OK", usr)
    }
    return CreateResponse("BADREQUEST", req)
}

func ValidateSession(req request.AuthenticateSession) (model.ServiceResponse) {
    session := GetSessionCache(req.SessionId)

    if (time.Now().Before(session.ExpiresAt)) {
        return CreateResponse("OK", session)
    }
    return CreateResponse("BADREQUEST", req)
}

