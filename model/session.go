package model

import (
    "github.com/google/uuid"
    "time"
)

type Session struct {
    ID              uuid.UUID
    UserId          uuid.UUID

    Permissions     []string

    ExpiresAt       time.Time
}

