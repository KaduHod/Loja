package main

import (
	"api-loja/src/infra/database"
	"api-loja/src/utils"
	"database/sql"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)
func GetBd() *sql.DB {
    con, err := database.NewConnection()
    if err != nil {
        log.Fatal()
    }
    return con
}
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("token")
        if token == "" {
            c.AbortWithStatus(401)
            return
        }
        body, err := json.Marshal(map[string]string{"token":token})

        if err != nil {
            c.AbortWithStatus(500)
            return
        }
        _, code, err := utils.Request(utils.RequestConfigInput{
            Method: "POST",
            Url: "http://loja-auth/verify-token",
            Data: string(body),
            DataType: "json",
        })
        if err != nil {
            c.AbortWithStatus(500)
            return
        }
        if code != 200 {
            c.AbortWithStatus(401)
            return
        }
        c.Next()
    }
}
func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loanding .env file")
    }
    server := gin.Default()
    server.Use(AuthMiddleware())
    server.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message":"ok",
        })
    })
    server.Run("0.0.0.0:8080")
}
