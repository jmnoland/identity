package model

import (
    "github.com/google/uuid"
    "time"
)

type Client struct {
    ID uuid.UUID
    Application string
    Name string

    CreatedAt time.Time
    ModifiedAt time.Time
}

