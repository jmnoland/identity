package service

import (
    "go.mongodb.org/mongo-driver/bson"
    "time"
    "errors"
    "github.com/jmnoland/identity/model"
    "github.com/jmnoland/identity/request"
    "github.com/jmnoland/identity/repository"
)

var (
    ErrInvalidEvent = errors.New("Invalid event")
)

func NewEvent(req request.EventRequest) (*model.Event, error) {
    isValid := validateEvent(req)

    if (!isValid) {
        return nil, ErrInvalidEvent;
    }

    event := model.Event{
        Application: req.Application,
        Action: req.Action,
        Type: req.Type,
        Request: req.Request,
        ActionRequestId: req.ActionRequestId,
        CreatedAt: time.Now(),
    }

    return &event, nil
}

func processClientEvent(event model.Event) {
    byteData, _ := bson.Marshal(event.Request)
    if event.Action == model.Actions["Create"] {
        var req request.CreateClientRequest
        bson.Unmarshal(byteData, &req)
        CreateClient(req, false)
    } else if event.Action == model.Actions["Update"] {
        var req request.UpdateClientRequest
        bson.Unmarshal(byteData, &req)
        UpdateClient(req, false)
    } else if event.Action == model.Actions["Delete"] {
        var req request.DeleteClientRequest
        bson.Unmarshal(byteData, &req)
        DeleteClient(req, false)
    }
}
func processCredentialEvent(event model.Event) {
    byteData, _ := bson.Marshal(event.Request)
    if event.Action == model.Actions["Create"] {
        var req request.CreateCredentialRequest
        bson.Unmarshal(byteData, &req)
        CreateCredential(req, false)
    }
}
func processUserEvent(event model.Event) {
    byteData, _ := bson.Marshal(event.Request)
    if event.Action == model.Actions["Create"] {
        var req request.CreateUserRequest
        bson.Unmarshal(byteData, &req)
        CreateUser(req, false)
    } else if event.Action == model.Actions["Update"] {
        var req request.UpdateUserRequest
        bson.Unmarshal(byteData, &req)
        UpdateUser(req, false)
    } else if event.Action == model.Actions["Delete"] {
        var req request.DeleteUserRequest
        bson.Unmarshal(byteData, &req)
        DeleteUser(req, false)
    }
}

func StartupProcessEvents() {
    events := repository.GetEvents()

    for _, event := range events {
        if event.Type == model.EventTypes["Client"] {
            processClientEvent(event);
        } else if event.Type == model.EventTypes["User"] {
            processUserEvent(event);
        } else if event.Type == model.EventTypes["Credential"] {
            processCredentialEvent(event);
        }
    }
}

func validateEvent(req request.EventRequest) (bool) {
    if req.Action == "" {
        return false;
    }
    if req.Type == "" {
        return false;
    }
    if req.Application == "" {
        return false;
    }

    return true;
}

