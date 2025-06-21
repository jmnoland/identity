package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/jmnoland/identity/request"
    "github.com/jmnoland/identity/service"
)

func CreateUser(c *gin.Context) {
    var newUser request.CreateUserRequest

    if err := c.BindJSON(&newUser); err != nil {
        return
    }

    response := service.CreateUser(newUser)

    SendResponse(c, response)
}

func CreateUserWithCredential(c *gin.Context) {
    var newUser request.CreateUserWithCredentialRequest

    if err := c.BindJSON(&newUser); err != nil {
        SendResponse(c, service.CreateResponse("BADREQUEST", err))
        return
    }

    response := service.CreateUserWithCredential(newUser)

    SendResponse(c, response)
}

func DeleteUser(c *gin.Context) {
    var removeRequest request.DeleteUserRequest

    if err := c.BindJSON(&removeRequest); err != nil {
        SendResponse(c, service.CreateResponse("BADREQUEST", err))
        return
    }

    response := service.DeleteUser(removeRequest)

    SendResponse(c, response)
}

func UpdateUser(c *gin.Context) {
    var updateRequest request.UpdateUserRequest

    if err := c.BindJSON(&updateRequest); err != nil {
        SendResponse(c, service.CreateResponse("BADREQUEST", err))
        return
    }

    response := service.UpdateUser(updateRequest)

    SendResponse(c, response)
}

func GetUser(c *gin.Context) {
    var getRequest request.GetUserRequest

    if err := c.BindJSON(&getRequest); err != nil {
        SendResponse(c, service.CreateResponse("BADREQUEST", err))
        return
    }

    client := service.GetUser(getRequest.ClientId, getRequest.UserId)
    response := service.CreateResponse("OK", client)

    SendResponse(c, response)
}

func GetUserByName(c *gin.Context) {
    var getRequest request.GetUserRequest

    if err := c.BindJSON(&getRequest); err != nil {
        SendResponse(c, service.CreateResponse("BADREQUEST", err))
        return
    }

    client := service.GetUserByName(getRequest.ClientId, getRequest.UserName)
    response := service.CreateResponse("OK", client)

    SendResponse(c, response)
}

