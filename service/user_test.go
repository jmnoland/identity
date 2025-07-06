package service

import (
    "testing"
    "time"
    "github.com/google/uuid"
    "github.com/jmnoland/identity/request"
    "github.com/jmnoland/identity/model"
)

const userIdStr string = "b6c649f4-2538-4344-93c6-2e2397b8eb9f";
const clientIdStr string = "485733f7-ff20-42b6-9bc6-c966e83cf314";
const applicationIdStr string = "8908242b-3995-4773-a097-22e977209974";

func TestCreateUser(t *testing.T) {
    userId := uuid.MustParse(userIdStr)
    clientId := uuid.MustParse(clientIdStr)
    application := model.Application{
        ID: uuid.MustParse(applicationIdStr),
        Name: "TestApp",
        Permissions: []string {"create"},
        ModifiedAt: time.Now(),
    }

    req := request.CreateUserWithCredentialRequest{
        User: request.CreateUserRequest{
            Application: []model.Application { application },
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
    result := CreateUserWithCredential(req, false)
    
    if result.ResponseCode == model.ResponseTypes["BADREQUEST"] {
        t.Error("Test Failed, %w", result.ResponseObject)
    }
}

func TestUpdateUser(t *testing.T) {
    userId := uuid.MustParse(userIdStr)
    clientId := uuid.MustParse(clientIdStr)
    application := model.Application{
        ID: uuid.MustParse(applicationIdStr),
        Name: "TestingApp",
        Permissions: []string {"create"},
        ModifiedAt: time.Now(),
    }

    req := request.UpdateUserRequest {
        UserId: userId,
        Application: []model.Application { application },
        RequestId: uuid.New(),
        ClientId: clientId,
        UserName: "test@email.co.za",
    }
    result := UpdateUser(req, false)
    
    if result.ResponseCode == model.ResponseTypes["BADREQUEST"] {
        t.Error("Test Failed, %w", result.ResponseObject)
    }

    updatedUser := result.ResponseObject.(model.User)
    if updatedUser.Name != req.UserName {
        t.Error("Test Failed, incorrect userName, %w", updatedUser.Name)
    }
}

func TestDeleteUser(t *testing.T) {
    userId := uuid.MustParse(userIdStr)
    clientId := uuid.MustParse(clientIdStr)
    application := model.Application{
        ID: uuid.MustParse(applicationIdStr),
        Name: "TestApp",
        Permissions: []string {"create"},
        ModifiedAt: time.Now(),
    }

    req := request.DeleteUserRequest{
        Application: []model.Application { application },
        RequestId: uuid.New(),
        UserId: userId,
        ClientId: clientId,
    }
    result := DeleteUser(req, false)
    
    if result.ResponseCode == model.ResponseTypes["BADREQUEST"] {
        t.Error("Test Failed, %w", result.ResponseObject)
    }

    user := GetUser(clientId, userId)
    if user.Name != "" {
        t.Error("Test Failed, user not deleted")
    }
}

