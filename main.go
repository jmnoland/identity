package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting")
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    fmt.Println("Running");
    r.Run()
}

