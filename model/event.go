package model

import (
    "github.com/google/uuid"
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Event struct {
    ID              primitive.ObjectID   `json:"id" bson:"_id,omitempty"`
    Application     string
    Action          string
    Type            string
    Request         string
    ActionRequestId uuid.UUID

    CreatedAt time.Time
}

