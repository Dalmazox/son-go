package middlewares

import (
	"github.com/dalmazox/son-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("request received", logrus.Fields{})

		c.Next()

		logger.Info("request finished", logrus.Fields{})
	}
}
