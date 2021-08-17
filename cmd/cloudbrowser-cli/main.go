package main

import (
	"github.com/hachi-n/cloudbrowser/cmd/cloudbrowser-cli/commands/ec2"
	"github.com/hachi-n/cloudbrowser/internal/config"
	"github.com/hachi-n/cloudbrowser/internal/envs"
	"github.com/hachi-n/cloudbrowser/internal/logger"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"os"
)

func init() {
	_log := logger.NewLogger()

	appEnv := envs.AppEnv()
	if appEnv == "" {
		_log.Fatal("APP_ENV error. please set APP_ENV=[development/staging/production]")
	}

	awsConfig := config.NewAwsConfig()
	_log.Debug("aws config values", zap.Object("awsConfig", awsConfig))

	err := awsConfig.Validate()
	if err != nil {
		_log.Fatal(err.Error(), zap.Object("awsConfig", awsConfig))
	}
}

func main() {
	_log := logger.NewLogger()
	app := &cli.App{
		Name:  "cloudbrowser-cli",
		Usage: "ec2 instance backup",
		Commands: []*cli.Command{
			ec2.Cmd(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		_log.Fatal("Fatal Error:", zap.String("err", err.Error()))
	}
}
