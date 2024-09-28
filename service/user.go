package service

import (
	"github.com/google/uuid"
	"github.com/jmnoland/identity/model"
	"github.com/jmnoland/identity/request"
    "github.com/jmnoland/identity/repository"
	"time"
)

func createUserEventRequest(req any, application string, action string, requestId uuid.UUID) request.EventRequest {
	eventReq := request.EventRequest{
		Application:     application,
		Type:            "User",
		Action:          action,
		ActionRequestId: requestId,
		Request:         req,
	}

	return eventReq
}

func CreateUser(req request.CreateUserRequest) model.ServiceResponse {
    existingUser := GetUserByName(req.ClientId, req.UserName)
    if existingUser.Name != "" {
        return CreateResponse("BADREQUEST", existingUser)
    }

	eventReq := createUserEventRequest(req, req.Application, model.Actions["Create"], req.RequestId)
	event, err := NewEvent(eventReq)
	if err != nil {
		panic(err)
	}

	client := model.User{
		ID:          req.UserId,
		Name:        req.UserName,
		Application: req.Application,
		CreatedAt:   time.Now(),
	}

	AddUserCache(client)

    repository.AddEvent(*event)

	return CreateResponse("CREATED", client)
}

func UpdateUser(req request.UpdateUserRequest) model.ServiceResponse {
	event := createUserEventRequest(req, req.Application, model.Actions["Update"], req.RequestId)
    _, err := NewEvent(event)
	if err != nil {
		panic(err)
	}

	client := GetUser(req.ClientId, req.UserId)
	client.Name = req.UserName
	client.ModifiedAt = time.Now()

	UpdateUserCache(client)

	return CreateResponse("UPDATED", client)
}

func DeleteUser(req request.DeleteUserRequest) model.ServiceResponse {
	event := createUserEventRequest(req, req.Application, model.Actions["Delete"], req.RequestId)
    _, err := NewEvent(event)
	if err != nil {
		panic(err)
	}

	client := GetUser(req.ClientId, req.UserId)

	RemoveUserCache(client)

    return CreateResponse("DELETED", nil)
}

