package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type ContentController struct {
	group *gin.RouterGroup
}

func NewContentController(group *gin.RouterGroup) *ContentController {
	return &ContentController{
		group: group,
	}
}

func (controller *ContentController) BindRoutes() {
	controller.group.GET("/", getContent)
}

func getContent(context *gin.Context) {
	context.JSON(200, gin.H{
		"content": viper.GetString("content"),
	})
}
