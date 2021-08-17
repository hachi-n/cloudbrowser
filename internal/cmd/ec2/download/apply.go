package download

import (
	"github.com/hachi-n/cloudbrowser/internal/config"
	form "github.com/hachi-n/cloudbrowser/internal/forms/ec2/download"
	"github.com/hachi-n/cloudbrowser/internal/models/downloader"
	"github.com/hachi-n/cloudbrowser/internal/models/downloader/ec2"
)

func Apply(form *form.DownloadForm) error {
	config.Aws.Override(form.Region)

	model := ec2.NewEc2Downloader()
	err := downloader.Download(model)
	if err != nil {
		return err
	}
	return nil
}
