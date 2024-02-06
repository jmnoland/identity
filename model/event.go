package model

import (
    "github.com/google/uuid"
    "time"
)

type Event struct {
    ID uuid.UUID
    CreatedAt time.Time
}

