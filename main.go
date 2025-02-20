package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    // Создаем новый роутер Gin
    router := gin.Default()

    // Определяем маршруты
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    // Запускаем сервер на порту 8080
    router.Run(":8080")
}