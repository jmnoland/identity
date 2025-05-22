package service

import (
    "time"
    "errors"
    "github.com/jmnoland/identity/model"    
    "github.com/jmnoland/identity/request"
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

