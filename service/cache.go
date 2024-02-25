package service

import (
    "github.com/google/uuid"
    "github.com/jmnoland/identity/model"
)

var clients = map[uuid.UUID]model.Client{}

func AddClient(client model.Client) {
    clients[client.ID] = client
}

