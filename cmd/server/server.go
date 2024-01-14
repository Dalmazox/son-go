package main

import (
	"github.com/dalmazox/son-go/cmd/server/middlewares"
	"github.com/dalmazox/son-go/config"
	"github.com/gin-gonic/gin"
)

func setupEngine() *gin.Engine {
	mode := gin.ReleaseMode

	if config.AppConfig.Environment == "development" {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)

	engine := gin.Default()

	engine.Use(
		middlewares.LogMiddleware(),
		gin.Recovery(),
	)

	engine.SetTrustedProxies(nil)

	bindRoutes(engine)

	return engine
}
