package request

import (
    "github.com/google/uuid"
)

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

