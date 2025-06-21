package service

import (
    "github.com/google/uuid"
    "github.com/jmnoland/identity/model"
)

var cache = model.Cache{
    Clients: map[uuid.UUID]model.Client{
        uuid.MustParse("00000000-0000-0000-0000-000000000000"): {
            ID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
        },
    },
    Credentials: map[string]model.Credential{
    },
    Sessions: map[uuid.UUID]model.Session{
    },
}

func GetCurrentCache() (model.Cache) {
    return cache
}

func AddSessionCache(session model.Session) {
    cache.Sessions[session.ID] = session;
}
func RemoveSessionCache(session model.Session) {
    delete(cache.Sessions, session.ID)
}
func GetSessionCache(id uuid.UUID) (model.Session) {
    return cache.Sessions[id]
}
func GetSessionForUserId(userId uuid.UUID) (model.Session) {
    for k, v := range cache.Sessions {
        if v.UserId == userId {
            return cache.Sessions[k]
        }
    }
    return model.Session{}
}

func AddClientCache(client model.Client) {
    cache.Clients[client.ID] = client;
}

func RemoveClientCache(client model.Client) {
    delete(cache.Clients, client.ID);
}

func UpdateClientCache(client model.Client) {
    cache.Clients[client.ID] = client;
}

func GetClient(id uuid.UUID) (model.Client) {
    return cache.Clients[id];
}

func GetClientByName(name string) (model.Client) {
    for k, v := range cache.Clients {
        if v.Name == name {
            return cache.Clients[k]
        }
    }
    return model.Client{}
}

func AddUserCache(user model.User) {
    client := cache.Clients[user.ClientId]
    client.Users = append(client.Users, user)
    cache.Clients[user.ClientId] = client
}

func RemoveUserCache(user model.User) {
    client := cache.Clients[user.ClientId]
    for i := range client.Users {
        if client.Users[i].ID == user.ID {
            client.Users = append(client.Users[:i], client.Users[i+1:]...)
        }
    }
    cache.Clients[user.ClientId] = client
}

func addUpdateApplicationToUser(exists model.User, update model.Application) {
    for i := range exists.Applications {
        if (exists.Applications[i].ID == update.ID) {
            exists.Applications[i].Name = update.Name
            exists.Applications[i].Permissions = append(exists.Applications[i].Permissions, update.Permissions...)
            return
        }
    }
    exists.Applications = append(exists.Applications, update)
}

func UpdateUserCache(user model.User) {
    client := cache.Clients[user.ClientId]
    usrs := client.Users
    for i := range usrs {
        if usrs[i].ID == user.ID {
            usrs[i].Name = user.Name
            addUpdateApplicationToUser(usrs[i], user.Applications[0])
        }
    }
    cache.Clients[user.ClientId] = client
}

func GetUser(clientId uuid.UUID, id uuid.UUID) (model.User) {
    for i := range cache.Clients[clientId].Users {
        if cache.Clients[clientId].Users[i].ID == id {
            return cache.Clients[clientId].Users[i]
        }
    }
    return model.User{}
}

func GetUserByName(clientId uuid.UUID, name string) (model.User) {
    client := cache.Clients[clientId]
    for i := range client.Users {
        if client.Users[i].Name == name {
            return client.Users[i]
        }
    }
    return model.User{}
}

func AddCredentialCache(credential model.Credential) {
    cache.Credentials[credential.Identifier] = credential
}

func GetCredential(identifier string) (model.Credential) {
    credential := cache.Credentials[identifier]
    return credential;
}

