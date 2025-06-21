package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/jmnoland/identity/service"
)

func GetCurrentCache(c *gin.Context) {
    cache := service.GetCurrentCache()
    response := service.CreateResponse("OK", cache)

    SendResponse(c, response)
}

