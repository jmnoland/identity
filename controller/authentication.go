package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/jmnoland/identity/request"
    "github.com/jmnoland/identity/service"
)

func AuthenticateSession(c *gin.Context) {
    var authRequest request.AuthenticateRequest

    if err := c.BindJSON(&authRequest); err != nil {
        SendResponse(c, service.CreateResponse("BADREQUEST", err))
        return
    }

    session := service.AuthenticateSession(authRequest)
    response := service.CreateResponse("OK", session)

    SendResponse(c, response)
}

func AuthenticateApiKey(c *gin.Context) {
    var authRequest request.AuthenticateRequest

    if err := c.BindJSON(&authRequest); err != nil {
        SendResponse(c, service.CreateResponse("BADREQUEST", err))
        return
    }

    user := service.AuthenticateApiKey(authRequest)
    response := service.CreateResponse("OK", user)

    SendResponse(c, response)
}

func ValidateSession(c *gin.Context) {
    var authRequest request.AuthenticateSession

    if err := c.BindJSON(&authRequest); err != nil {
        SendResponse(c, service.CreateResponse("BADREQUEST", err))
        return
    }

    session := service.ValidateSession(authRequest)
    response := service.CreateResponse("OK", session)

    SendResponse(c, response)
}

