package logger

import (
	"fmt"
	"github.com/hachi-n/cloudbrowser/internal/envs"
	"go.uber.org/zap"
	"os"
)

var loggerCache *zap.Logger

func NewLogger() (logger *zap.Logger) {
	if loggerCache != nil {
		return loggerCache
	}

	var err error

	switch envs.AppEnv() {
	case "production":
		logger, err = zap.NewProduction()
	case "staging":
		logger, err = zap.NewProduction()
	case "development":
		logger, err = zap.NewDevelopment()
	default:
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		fmt.Println("Un expected error: check your logger.")
		os.Exit(1)
	}

	loggerCache = logger
	return
}
