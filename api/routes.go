package main

import (
	"logger-service/api/handlers"

	"github.com/gin-gonic/gin"
)

func routes() *gin.Engine {
	router := gin.Default()

	router.GET("/", handlers.Home)
	router.GET("/favicon.ico", handlers.DoNothing)

	return router
}
