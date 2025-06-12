package request

import (
    "github.com/google/uuid"
    "github.com/jmnoland/identity/model"
)

type CreateUserWithCredentialRequest struct {
    User            CreateUserRequest
    Credential      CreateCredentialRequest
}

type CreateUserRequest struct {
    Application     []model.Application
    RequestId       uuid.UUID

    UserId          uuid.UUID
    UserName        string
    ClientId        uuid.UUID
}

type UpdateUserRequest struct {
    Application     []model.Application
    RequestId       uuid.UUID

    UserId          uuid.UUID
    UserName        string
    ClientId        uuid.UUID
}

type DeleteUserRequest struct {
    Application     []model.Application
    RequestId       uuid.UUID

    UserId          uuid.UUID
    ClientId        uuid.UUID
}

type GetUserRequest struct {
    UserId          uuid.UUID
    UserName        string

    ClientId        uuid.UUID
}

