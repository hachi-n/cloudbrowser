package config

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/goccy/go-yaml"
	"github.com/hachi-n/cloudbrowser/internal/envs"
	"github.com/hachi-n/cloudbrowser/internal/logger"
	_ "github.com/hachi-n/cloudbrowser/pack"
	"github.com/rakyll/statik/fs"
	"go.uber.org/zap/zapcore"
	"os"
)

type AwsConfig struct {
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
	Region    string `yaml:"region"`
}

var (
	Aws *AwsConfig
)

func (c *AwsConfig) Override(region string) {
	if region == "" {
		return
	}
	c.Region = region
}

func (c *AwsConfig) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("aws_access_key", c.AccessKey)
	enc.AddString("aws_secret_key", c.SecretKey)
	enc.AddString("aws_region", c.Region)
	return nil
}

func (c *AwsConfig) Validate() error {
	return validation.ValidateStruct(c,
		// AccessKey
		validation.Field(&c.AccessKey, validation.Required),
		// SecretKey
		validation.Field(&c.SecretKey, validation.Required),
	)
}

const (
	AWS_CONFIG_PATH = "/configs/aws.yaml"
)

const (
	DATA_DIR = "./data"
)

var _log = logger.NewLogger()

func NewAwsConfig() *AwsConfig {
	if Aws != nil {
		return Aws
	}

	awsConfig := getAwsConfigFromEnv()
	if awsConfig != nil {
		return awsConfig
	}

	awsConfig, err := getAwsConfigFromYaml()
	if err != nil {
		_log.Fatal(err.Error())
	}

	Aws = awsConfig
	return awsConfig
}

func getAwsConfigFromYaml() (*AwsConfig, error) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	f, err := statikFS.Open(AWS_CONFIG_PATH)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	yamlContentPath, err := yaml.PathString(
		fmt.Sprintf("$.%s", envs.AppEnv()),
	)
	if err != nil {
		return nil, err
	}

	awsConfig := new(AwsConfig)
	if err := yamlContentPath.Read(f, awsConfig); err != nil {
		fmt.Println("yaml load err.", err)
		return nil, err
	}
	return awsConfig, nil
}

func getAwsConfigFromEnv() *AwsConfig {
	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")
	region := os.Getenv("AWS_REGION")

	if accessKey == "" || secretKey == "" {
		return nil
	}

	return &AwsConfig{
		AccessKey: accessKey,
		SecretKey: secretKey,
		Region:    region,
	}
}
