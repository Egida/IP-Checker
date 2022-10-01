package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func InitRouter() {
	godotenv.Load()
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/"+os.Getenv("REQUESTTOKEN")+"/:ip", IPCheck)
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
