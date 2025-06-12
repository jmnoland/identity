package request

import (
    "github.com/google/uuid"
)

type AuthenticateRequest struct {
    Application     string
    RequestId       uuid.UUID

    Identifier      string
    Secret          string

    ClientId        uuid.UUID
    ClientName      string
}

type AuthenticateSession struct {
    RequestId       uuid.UUID

    SessionId       uuid.UUID
}

