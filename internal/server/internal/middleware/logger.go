package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hachi-n/cloudbrowser/internal/logger"
	"go.uber.org/zap"
	"time"
)

var _log = logger.NewLogger()

func ServerLogFormat(c *gin.Context) {
	oldTime := time.Now()
	userAgent := c.GetHeader("User-Agent")
	c.Next()
	_log.Info("incoming request",
		zap.String("path", c.Request.URL.Path),
		zap.String("User-Agent", userAgent),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("elapsed", time.Now().Sub(oldTime)),
	)
}
