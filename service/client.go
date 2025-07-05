package service

import (
	"github.com/google/uuid"
	"github.com/jmnoland/identity/model"
	"github.com/jmnoland/identity/request"
    "github.com/jmnoland/identity/repository"
	"time"
)

func createClientEventRequest(req any, application string, action string, requestId uuid.UUID) request.EventRequest {
	eventReq := request.EventRequest{
		Application:     application,
		Type:            model.EventTypes["Client"],
		Action:          action,
		ActionRequestId: requestId,
		Request:         req,
	}

	return eventReq
}

func CreateClient(req request.CreateClientRequest, createEvent bool) model.ServiceResponse {
    existingClient := GetClientByName(req.ClientName)
    if existingClient.Name != "" {
        return CreateResponse("BADREQUEST", existingClient)
    }

	eventReq := createClientEventRequest(req, req.Application, model.Actions["Create"], req.RequestId)
	event, err := NewEvent(eventReq)
	if err != nil {
		panic(err)
	}

	client := model.Client{
		ID:          req.ClientId,
		Name:        req.ClientName,
		Application: req.Application,
		CreatedAt:   time.Now(),
	}

	AddClientCache(client)

    if createEvent {
        repository.AddEvent(*event)
    }

	return CreateResponse("CREATED", client)
}

func UpdateClient(req request.UpdateClientRequest, createEvent bool) model.ServiceResponse {
	eventReq := createClientEventRequest(req, req.Application, model.Actions["Update"], req.RequestId)
    event, err := NewEvent(eventReq)
	if err != nil {
		panic(err)
	}

	client := GetClient(req.ClientId)
	client.Name = req.ClientName
	client.ModifiedAt = time.Now()

	UpdateClientCache(client)

    if createEvent {
        repository.AddEvent(*event)
    }

	return CreateResponse("UPDATED", client)
}

func DeleteClient(req request.DeleteClientRequest, createEvent bool) model.ServiceResponse {
	eventReq := createClientEventRequest(req, req.Application, model.Actions["Delete"], req.RequestId)
    event, err := NewEvent(eventReq)
	if err != nil {
		panic(err)
	}

	client := GetClient(req.ClientId)

	RemoveClientCache(client)

    if createEvent {
        repository.AddEvent(*event)
    }

    return CreateResponse("DELETED", nil)
}
