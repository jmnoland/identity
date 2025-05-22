package request

import (
    "github.com/google/uuid"
)

type CreateUserWithCredentialRequest struct {
    User            CreateUserRequest
    Credential      CreateCredentialRequest
}

type CreateUserRequest struct {
    Application     string
    RequestId       uuid.UUID

    UserId          uuid.UUID
    UserName        string
    ClientId        uuid.UUID
}

type UpdateUserRequest struct {
    Application     string
    RequestId       uuid.UUID

    UserId          uuid.UUID
    UserName        string
    ClientId        uuid.UUID
}

type DeleteUserRequest struct {
    Application     string
    RequestId       uuid.UUID

    UserId          uuid.UUID
    ClientId        uuid.UUID
}

type GetUserRequest struct {
    UserId          uuid.UUID
    UserName        string

    ClientId        uuid.UUID
}

