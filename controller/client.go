package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/jmnoland/identity/request"
    "github.com/jmnoland/identity/service"
)

func CreateClient(c *gin.Context) {
    var newClient request.CreateClientRequest

    if err := c.BindJSON(&newClient); err != nil {
        return
    }

    response := service.CreateClient(newClient)

    SendResponse(c, response)
}

func DeleteClient(c *gin.Context) {
    var removeRequest request.DeleteClientRequest

    if err := c.BindJSON(&removeRequest); err != nil {
        return
    }

    response := service.DeleteClient(removeRequest)
    
    SendResponse(c, response)
}

func UpdateClient(c *gin.Context) {
    var updateRequest request.UpdateClientRequest

    if err := c.BindJSON(&updateRequest); err != nil {
        return
    }

    response := service.UpdateClient(updateRequest)

    SendResponse(c, response)
}

func GetClient(c *gin.Context) {
    var getRequest request.GetClientRequest

    if err := c.BindJSON(&getRequest); err != nil {
        return
    }

    client := service.GetClient(getRequest.ClientId)
    response := service.CreateResponse("FOUND", client)

    SendResponse(c, response)
}

func GetClientByName(c *gin.Context) {
    var getRequest request.GetClientRequest

    if err := c.BindJSON(&getRequest); err != nil {
        return
    }

    client := service.GetClientByName(getRequest.ClientName)
    response := service.CreateResponse("FOUND", client)

    SendResponse(c, response)
}

