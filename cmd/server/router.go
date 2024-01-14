package main

import (
	"github.com/dalmazox/son-go/cmd/server/controllers"
	"github.com/gin-gonic/gin"
)

func bindRoutes(e *gin.Engine) {
	healthController := controllers.NewHealthController()
	contentController := controllers.NewContentController(e.Group("v1/content"))

	e.GET("/health", healthController.Health)
	contentController.BindRoutes()
}
