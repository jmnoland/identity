package service

import (
    "github.com/google/uuid"
    "github.com/jmnoland/identity/model"
)

var clients = map[uuid.UUID]model.Client{}

func AddClientCache(client model.Client) {
    clients[client.ID] = client;
}

func RemoveClientCache(client model.Client) {
    delete(clients, client.ID);
}

func UpdateClientCache(client model.Client) {
    clients[client.ID] = client;
}

func GetClient(id uuid.UUID) (model.Client) {
    return clients[id]
}

