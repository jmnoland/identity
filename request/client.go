package request

import (
    "github.com/google/uuid"
)

type CreateClientRequest struct {
    Application string
    RequestId uuid.UUID

    ClientId uuid.UUID
    ClientName string
}

