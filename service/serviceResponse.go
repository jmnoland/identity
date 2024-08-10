package service

import (
    "github.com/jmnoland/identity/model"
)

func CreateResponse(responseType string, responseObject any) (model.ServiceResponse) {
    return model.ServiceResponse{
        ResponseCode: model.ResponseTypes[responseType],
        ResponseObject: responseObject,
    }
}

