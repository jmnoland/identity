package service

import (
    "testing"
    "fmt"
    "github.com/google/uuid"
    "github.com/jmnoland/identity/request"
)

func TestNewEvent(t *testing.T) {
    event := request.EventRequest{
        Application: "TestApp",
        Action: "Create",
        Type: "User",
        Request: "",
        ActionRequestId: uuid.New(),
    }
    _, err := NewEvent(event)
    
    if err != nil {
        t.Error("Test Failed,", fmt.Errorf("error: %w", err))
    }
}

func TestNewEventFailValidation(t *testing.T) {
    var tests = []struct {
        name string
        val request.EventRequest
    }{
        { "Empty Application", request.EventRequest{
            Application: "",
            Action: "Create",
            Type: "User",
            Request: "",
            ActionRequestId: uuid.New(),
        }},
        { "Empty Action", request.EventRequest{
            Application: "TestApp",
            Action: "",
            Type: "User",
            Request: "",
            ActionRequestId: uuid.New(),
        }},
        { "Empty Type", request.EventRequest{
            Application: "TestApp",
            Action: "Create",
            Type: "",
            Request: "",
            ActionRequestId: uuid.New(),
        }},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            event := request.EventRequest{
                Application: "",
                Action: "",
                Type: "",
                Request: "",
                ActionRequestId: uuid.New(),
            }
            result, _ := NewEvent(event)
            
            if result != nil {
                t.Error("Test Failed, %w", result)
            }
        })
    }
}

