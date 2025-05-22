package service

import (
    "testing"
    "github.com/google/uuid"
    "github.com/jmnoland/identity/request"
    "github.com/jmnoland/identity/model"
)

func TestCreateUser(t *testing.T) {
    userId := uuid.New()
    clientId := uuid.New()
    req := request.CreateUserWithCredentialRequest{
        User: request.CreateUserRequest{
            Application: "TestApp",
            RequestId: uuid.New(),
            UserId: userId,
            UserName: "test@email.com",
            ClientId: clientId,
        },
        Credential: request.CreateCredentialRequest{
            UserId: userId,
            ClientId: clientId,
            Application: "TestApp",
            Type: "Password",
            Identifier: "test@email.com",
            Secret: "Password1234",
            RequestId: uuid.New(),
        },
    }
    result := CreateUserWithCredential(req)
    
    if result.ResponseCode == model.ResponseTypes["BADREQUEST"] {
        t.Error("Test Failed, %w", result.ResponseObject)
    }
}

