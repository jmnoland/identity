package service

import (
    "os"
    "strconv"
    "crypto/rand"
    "encoding/base64"
    "golang.org/x/crypto/argon2"
	"github.com/google/uuid"
	"github.com/jmnoland/identity/model"
	"github.com/jmnoland/identity/request"
    "github.com/jmnoland/identity/repository"
)

var iterationsStr = os.Getenv("IDENTITY_ITERATIONS")
var parallelismStr = os.Getenv("IDENTITY_PARALLELISM")
var keyLenStr = os.Getenv("IDENTITY_KEYLEN")
var memoryStr = os.Getenv("IDENTITY_MEMORY")

func createCredentialEventRequest(req any, app string, action string, requestId uuid.UUID) request.EventRequest {
	eventReq := request.EventRequest{
		Application:     app,
		Type:            "Credential",
		Action:          action,
		ActionRequestId: requestId,
		Request:         req,
	}

	return eventReq
}

func CreateCredential(req request.CreateCredentialRequest) model.ServiceResponse {
    existingUser := GetUser(req.ClientId, req.UserId)
    if existingUser.Name == "" {
        return CreateResponse("BADREQUEST", existingUser)
    }

    hash, err := generateSecretFromString(req.Secret)
    if err != nil {
        return CreateResponse("ERROR", err)
    }

    req.Secret = hash

    credential := model.Credential{
        ID:         uuid.New(),
        UserId:     existingUser.ID,
        ClientId:   req.ClientId,
        Type:       req.Type,
        Identifier: req.Identifier,
        Secret:     hash,
    }

	eventReq := createCredentialEventRequest(req, req.Application, model.Actions["Create"], req.RequestId)
	event, err := NewEvent(eventReq)
	if err != nil {
		panic(err)
	}

	AddCredentialCache(credential)

    repository.AddEvent(*event)

	return CreateResponse("CREATED", credential)
}

func CreateSecretHashFromExisting(secret string, salt string) (str string) {
    iterations, _ := strconv.ParseInt(iterationsStr, 10, 32);
    memory, _ := strconv.ParseInt(memoryStr, 10, 32);
    parallelism, _ := strconv.ParseInt(parallelismStr, 10, 8);
    keyLen, _ := strconv.ParseInt(keyLenStr, 10, 32);

    saltString, _ := base64.RawStdEncoding.DecodeString(salt)
    hash := argon2.IDKey([]byte(secret), saltString, uint32(iterations), uint32(memory), uint8(parallelism), uint32(keyLen))

    hashString := base64.RawStdEncoding.EncodeToString(hash);
    return hashString
}

func generateSecretFromString(secret string) (str string, err error) {
    salt, err := generateRandomSalt(16)
    if err != nil {
        return "", err
    }

    iterations, _ := strconv.ParseInt(iterationsStr, 10, 32);
    memory, _ := strconv.ParseInt(memoryStr, 10, 32);
    parallelism, _ := strconv.ParseInt(parallelismStr, 10, 32);
    keyLen, _ := strconv.ParseInt(keyLenStr, 10, 32);
    hash := argon2.IDKey([]byte(secret), salt, uint32(iterations), uint32(memory), uint8(parallelism), uint32(keyLen))

    hashString := base64.RawStdEncoding.EncodeToString(hash)
    saltString := base64.RawStdEncoding.EncodeToString(salt)

    return hashString + ":" + saltString, nil
}

func generateRandomSalt(n uint32) ([]byte, error) {
    b := make([]byte, n)
    _, err := rand.Read(b)
    if err != nil {
        return nil, err
    }
    return b, nil
}

