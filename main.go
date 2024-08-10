package main

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/jmnoland/identity/controller"
)

func main() {
	fmt.Println("Starting")
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.POST("/client", controller.CreateClient)
    fmt.Println("Running");

    r.Run()
}

