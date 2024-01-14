package controllers

import "github.com/gin-gonic/gin"

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (controller *HealthController) Health(context *gin.Context) {
	context.JSON(200, gin.H{
		"status": "healthy",
	})
}
