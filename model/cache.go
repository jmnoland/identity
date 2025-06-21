package model

import (
    "github.com/google/uuid"
)

type Cache struct {
    Clients         map[uuid.UUID]Client
    Credentials     map[string]Credential
    Sessions        map[uuid.UUID]Session
}

