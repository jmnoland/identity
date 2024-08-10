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

