package model

import (
    "github.com/google/uuid"
    "time"
)

type Event struct {
    ID uuid.UUID
    Application string
    Action string
    Type string
    Request string
    ActionRequestId uuid.UUID

    CreatedAt time.Time
}

