package main

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "github.com/jmnoland/identity/controller"
    "github.com/jmnoland/identity/service"
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
    r.POST("/client/update", controller.UpdateClient)
    r.DELETE("/client", controller.DeleteClient)
    r.GET("/client", controller.GetClient)
    r.GET("/clientName", controller.GetClientByName)
    
    r.POST("/user", controller.CreateUser)
    r.POST("/user/credential", controller.CreateUserWithCredential)
    r.POST("/user/update", controller.UpdateUser)
    r.DELETE("/user", controller.DeleteUser)
    r.GET("/user", controller.GetUser)
    r.GET("/userName", controller.GetUserByName)

    r.POST("/auth/session", controller.AuthenticateSession)
    r.POST("/auth/key", controller.AuthenticateApiKey)
    r.POST("/auth/session/validate", controller.ValidateSession)

    r.GET("/cache", controller.GetCurrentCache)

    fmt.Println("Event processing")

    service.StartupProcessEvents()

    fmt.Println("Running")

    r.Run()
}

