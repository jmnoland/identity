package service

import (
    "time"
    "encoding/json"
    "github.com/jmnoland/identity/model"    
    "github.com/jmnoland/identity/request"
)

func createEventRequest(req request.CreateClientRequest) (request.EventRequest) {
    reqString, err := json.Marshal(req)
    if err != nil {
        panic(err)
    }

    eventReq := request.EventRequest{
        Application: req.Application,
        Type: "Client",
        Action: model.Actions["Create"],
        ActionRequestId: req.RequestId,
        Request: string(reqString),
    }

    return eventReq
}

func CreateClient(req request.CreateClientRequest) (model.Client) {
    event := createEventRequest(req)
    _, err := NewEvent(event)
    if err != nil {
        panic(err)
    }

    client := model.Client{
        ID: req.ClientId,
        Name: req.ClientName,
        Application: req.Application,
        CreatedAt: time.Now(),
    }

    AddClient(client)

    return client
}

