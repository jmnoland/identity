package model

import (
    "github.com/google/uuid"
    "time"
)

type User struct {
    ID              uuid.UUID
    ClientId        uuid.UUID
    Application     string
    Name            string

    CreatedAt       time.Time
    ModifiedAt      time.Time
}

