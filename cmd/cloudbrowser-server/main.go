package main

import (
	"github.com/hachi-n/cloudbrowser/internal/logger"
	server "github.com/hachi-n/cloudbrowser/internal/server/cloudbrowser-server"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
)

func main() {
	_log := logger.NewLogger()
	app := &cli.App{
		Name:   "cloudbrowser-server",
		Usage:  "cloudbrowser start",
		Action: serve,
	}

	err := app.Run(os.Args)
	if err != nil {
		_log.Fatal("Fatal Error:", zap.String("err", err.Error()))
	}
}

func serve(c *cli.Context) error {
	return server.StartDaemon()
}
