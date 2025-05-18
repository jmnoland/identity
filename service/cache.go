package service

import (
    "github.com/google/uuid"
    "github.com/jmnoland/identity/model"
)

var clients = map[uuid.UUID]model.Client{
    uuid.MustParse("00000000-0000-0000-0000-000000000000"): {
        ID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
    },
}

var credentials = map[string]model.Credential{

}

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
    return clients[id];
}

func GetClientByName(name string) (model.Client) {
    for k, v := range clients {
        if v.Name == name {
            return clients[k]
        }
    }
    return model.Client{}
}

func AddUserCache(user model.User) {
    client := clients[user.ClientId]
    client.Users = append(client.Users, user)
    clients[user.ClientId] = client
}

func RemoveUserCache(user model.User) {
    client := clients[user.ClientId]
    for i := range client.Users {
        if client.Users[i].ID == user.ID {
            client.Users = append(client.Users[:i], client.Users[i+1:]...)
        }
    }
    clients[user.ClientId] = client
}

func UpdateUserCache(user model.User) {
    client := clients[user.ClientId]
    usrs := client.Users
    for i := range usrs {
        if usrs[i].ID == user.ID {
            usrs[i].Name = user.Name
        }
    }
    clients[user.ClientId] = client
}

func GetUser(clientId uuid.UUID, id uuid.UUID) (model.User) {
    for i := range clients[clientId].Users {
        if clients[clientId].Users[i].ID == id {
            return clients[clientId].Users[i]
        }
    }
    return model.User{}
}

func GetUserByName(clientId uuid.UUID, name string) (model.User) {
    client := clients[clientId]
    for i := range client.Users {
        if client.Users[i].Name == name {
            return client.Users[i]
        }
    }
    return model.User{}
}

func AddCredentialCache(credential model.Credential) {
    credentials[credential.Identifier] = credential
}

func GetCredential(identifier string) (model.Credential) {
    credential := credentials[identifier]
    return credential;
}

