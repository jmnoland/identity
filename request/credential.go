package request

import (
    "github.com/google/uuid"
)

type CreateCredentialRequest struct {
    UserId          uuid.UUID
    ClientId        uuid.UUID
    Type            string
    Identifier      string
    Secret          string
    RequestId       uuid.UUID
}

