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
        usr := GetUser(cred.ClientId, cred.UserId)

        hoursToAdd, _ := strconv.ParseInt(expireHours, 10, 32)
        expiresAt := time.Now().Add(time.Hour * time.Duration(hoursToAdd))

        sess := model.Session{
            ID:             uuid.New(),
            Permissions:    findUserPermissions(usr.Applications, req.Application),
            ExpiresAt:      expiresAt,
        }
        AddSessionCache(sess)

        return CreateResponse("CREATED", sess)
    }
    return CreateResponse("BADREQUEST", req)
}

