package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
    fmt.Println("Hello")
    server := gin.Default()
    server.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message":"ok",
        })
    })
    server.Run("0.0.0.0:8080")
}
