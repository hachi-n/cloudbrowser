package ec2

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	service "github.com/hachi-n/cloudbrowser/internal/aws/services/ec2"
	"github.com/hachi-n/cloudbrowser/internal/config"
	"github.com/hachi-n/cloudbrowser/internal/logger"
	"github.com/hachi-n/cloudbrowser/internal/utils"
	"go.uber.org/zap"
	"path/filepath"
)

type Ec2Downloader struct {
	client *service.Client
}

const dataFileName = "ec2.json"

var _log = logger.NewLogger()

func NewEc2Downloader() *Ec2Downloader {
	return &Ec2Downloader{
		client: service.NewClient(),
	}
}

func (m *Ec2Downloader) Download(ctx context.Context) error {
	b, err := m.describeInstance(ctx)
	if err != nil {
		return err
	}

	dist := filepath.Join(config.DATA_DIR, dataFileName)
	if err = utils.ForceWriteFile(dist, b); err != nil {
		return err
	}

	_log.Info("download ok.", zap.String("distinarionPath", dist))
	return nil
}

func (m *Ec2Downloader) describeInstance(ctx context.Context) ([]byte, error) {
	input := &ec2.DescribeInstancesInput{}
	output, err := m.client.DescribeInstances(ctx, input)
	if err != nil {
		return nil, err
	}

	b, err := utils.PrettyJson(output)
	if err != nil {
		return nil, err
	}
	return b, err
}
