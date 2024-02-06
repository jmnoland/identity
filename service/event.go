package service

import (
    "errors"
    "github.com/jmnoland/identity/model"    
)

var (
    ErrInvalidEvent = errors.New("Invalid event")
)

func NewEvent() (Event, error) {
    
    return Event{
        
    }, nil
}

