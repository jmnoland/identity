package controller

import (
    "github.com/gin-gonic/gin"
    "github.com/jmnoland/identity/model"
)

func SendResponse(c *gin.Context, response model.ServiceResponse) {

    c.IndentedJSON(response.ResponseCode, response.ResponseObject)
}

