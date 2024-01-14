package main

import (
	"github.com/dalmazox/son-go/config"
	"github.com/dalmazox/son-go/pkg/logger"
	"github.com/sirupsen/logrus"
)

func init() {
	fields := logrus.Fields{"category": "configuration"}
	if err := config.LoadConfig(); err != nil {
		logger.Fatal(err.Error(), fields)
	}

	logger.Info("configuration loaded", fields)
}

func main() {
	logger.Info("main", logrus.Fields{"category": "main"})

	router := setupEngine()

	router.Run(":" + config.AppConfig.Port)
}
