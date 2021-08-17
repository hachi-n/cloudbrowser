package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/hachi-n/cloudbrowser/internal/aws/configs"
)

type Client = ec2.Client

var client *Client

func NewClient() *Client {
	if client != nil {
		return client
	}

	cfg := configs.NewConfig()

	return ec2.NewFromConfig(*cfg)
}
