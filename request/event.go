package request

import (
    "github.com/google/uuid"
)

type EventRequest struct {
    Application     string
    Action          string
    Type            string
    Request         any
    ActionRequestId uuid.UUID
}

