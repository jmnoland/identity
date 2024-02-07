package service

import (
    "errors"
    "github.com/google/uuid"
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
        ID: uuid.New(),
    }

    return &event, nil
}

func validateEvent(req request.EventRequest) (bool) {
    if req.Action != "" {
        return true;
    }

    return false;
}

