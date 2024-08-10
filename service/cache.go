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

func GetClientByName(name string) (model.Client) {
    for k, v := range clients {
        if v.Name == name {
            return clients[k]
        }
    }
    return model.Client{}
}

