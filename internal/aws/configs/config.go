package configs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	sdkconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/hachi-n/cloudbrowser/internal/aws/credentials"
	"github.com/hachi-n/cloudbrowser/internal/config"
	"github.com/hachi-n/cloudbrowser/internal/logger"
)

var cfg *aws.Config
var _log = logger.NewLogger()

func NewConfig() *aws.Config {
	if cfg != nil {
		return cfg
	}

	cred := credentials.NewFromKeys(
		config.Aws.AccessKey,
		config.Aws.SecretKey,
		"",
	)

	c, err := sdkconfig.LoadDefaultConfig(
		context.Background(),
		sdkconfig.WithCredentialsProvider(cred),
		sdkconfig.WithRegion(config.Aws.Region),
	)

	if err != nil {
		_log.Fatal("aws config not found")
	}

	cfg = &c
	return cfg
}
