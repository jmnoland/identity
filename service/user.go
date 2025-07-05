package service

import (
	"github.com/google/uuid"
	"github.com/jmnoland/identity/model"
	"github.com/jmnoland/identity/request"
    "github.com/jmnoland/identity/repository"
	"time"
)

func createUserEventRequest(req any, application model.Application, action string, requestId uuid.UUID) request.EventRequest {
	eventReq := request.EventRequest{
		Application:     application.Name,
		Type:            model.EventTypes["User"],
		Action:          action,
		ActionRequestId: requestId,
		Request:         req,
	}

	return eventReq
}

func CreateUserWithCredential(req request.CreateUserWithCredentialRequest, createEvent bool) model.ServiceResponse {
    existingUser := GetUserByName(req.User.ClientId, req.User.UserName)
    if existingUser.Name != "" {
        return CreateResponse("BADREQUEST", existingUser)
    }

	user := model.User{
		ID:          req.User.UserId,
		Name:        req.User.UserName,
		Applications:req.User.Application,
        ClientId:    req.User.ClientId,
		CreatedAt:   time.Now(),
	}

	AddUserCache(user)

	eventReq := createUserEventRequest(req.User, req.User.Application[0], model.Actions["Create"], req.User.RequestId)
	event, err := NewEvent(eventReq)
	if err != nil {
		panic(err)
	}

    if createEvent {
        repository.AddEvent(*event)
    }

    CreateCredential(req.Credential, createEvent)

	return CreateResponse("CREATED", user)
}

func CreateUser(req request.CreateUserRequest, createEvent bool) model.ServiceResponse {
    existingUser := GetUserByName(req.ClientId, req.UserName)
    if existingUser.Name != "" {
        return CreateResponse("BADREQUEST", existingUser)
    }

	user := model.User{
		ID:          req.UserId,
		Name:        req.UserName,
		Applications:req.Application,
        ClientId:    req.ClientId,
		CreatedAt:   time.Now(),
	}

	AddUserCache(user)

	eventReq := createUserEventRequest(req, req.Application[0], model.Actions["Create"], req.RequestId)
	event, err := NewEvent(eventReq)
	if err != nil {
		panic(err)
	}

    if createEvent {
        repository.AddEvent(*event)
    }

	return CreateResponse("CREATED", user)
}

func UpdateUser(req request.UpdateUserRequest, createEvent bool) model.ServiceResponse {
	eventReq := createUserEventRequest(req, req.Application[0], model.Actions["Update"], req.RequestId)
    event, err := NewEvent(eventReq)
	if err != nil {
		panic(err)
	}

	user := GetUser(req.ClientId, req.UserId)
	user.Name = req.UserName
	user.ModifiedAt = time.Now()

	UpdateUserCache(user)

    if createEvent {
        repository.AddEvent(*event)
    }

	return CreateResponse("UPDATED", user)
}

func DeleteUser(req request.DeleteUserRequest, createEvent bool) model.ServiceResponse {
	eventReq := createUserEventRequest(req, req.Application[0], model.Actions["Delete"], req.RequestId)
    event, err := NewEvent(eventReq)
	if err != nil {
		panic(err)
	}

	user := GetUser(req.ClientId, req.UserId)

	RemoveUserCache(user)

    if createEvent {
        repository.AddEvent(*event)
    }

    return CreateResponse("DELETED", nil)
}

