package model

import (
    "github.com/google/uuid"
    "time"
)

type Application struct {
    ID              uuid.UUID
    Name            string

    Permissions     []string

    ModifiedAt      time.Time
}

