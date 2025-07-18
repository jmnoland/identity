package model

import (
    "github.com/google/uuid"
)

type Credential struct {
    ID              uuid.UUID

    UserId          uuid.UUID
    ClientId        uuid.UUID
    Type            string
    Identifier      string
    Secret          string
}

