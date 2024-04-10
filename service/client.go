package service

import (
    "time"
    "encoding/json"
    "github.com/google/uuid"
    "github.com/jmnoland/identity/model"    
    "github.com/jmnoland/identity/request"
)

func createEventRequest(reqString []byte, application string, action string, requestId uuid.UUID) (request.EventRequest) {
    eventReq := request.EventRequest{
        Application: application,
        Type: "Client",
        Action: action,
        ActionRequestId: requestId,
        Request: string(reqString),
    }

    return eventReq
}

func CreateClient(req request.CreateClientRequest) (model.Client) {
    reqString, err := json.Marshal(req)
    if err != nil {
        panic(err)
    }

    event := createEventRequest(reqString, req.Application, model.Actions["Create"], req.RequestId)
    _, err = NewEvent(event)
    if err != nil {
        panic(err)
    }

    client := model.Client{
        ID: req.ClientId,
        Name: req.ClientName,
        Application: req.Application,
        CreatedAt: time.Now(),
    }

    AddClientCache(client)

    return client
}

func UpdateClient(req request.UpdateClientRequest) (model.Client) {
    reqString, err := json.Marshal(req)
    if err != nil {
        panic(err)
    }

    event := createEventRequest(reqString, req.Application, model.Actions["Create"], req.RequestId)
    _, err = NewEvent(event)
    if err != nil {
        panic(err)
    }

    client := GetClient(req.ClientId)
    client.Name = req.ClientName
    client.ModifiedAt = time.Now()

    UpdateClientCache(client)

    return client
}

func DeleteClient(req request.DeleteClientRequest) {
    reqString, err := json.Marshal(req)
    if err != nil {
        panic(err)
    }

    event := createEventRequest(reqString, req.Application, model.Actions["Delete"], req.RequestId)
    _, err = NewEvent(event)
    if err != nil {
        panic(err)
    }

    client := GetClient(req.ClientId)

    RemoveClientCache(client)
}

