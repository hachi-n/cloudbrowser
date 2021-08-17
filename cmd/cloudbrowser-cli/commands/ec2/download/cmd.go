package download

import (
	"github.com/hachi-n/cloudbrowser/cmd/cloudbrowser-cli/internal/flags"
	cmd "github.com/hachi-n/cloudbrowser/internal/cmd/ec2/download"
	form "github.com/hachi-n/cloudbrowser/internal/forms/ec2/download"
	"github.com/hachi-n/cloudbrowser/internal/logger"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var _log = logger.NewLogger()

func Cmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "download",
		Usage: "download ec2 ami",
		Flags: []cli.Flag{
			flags.RegionFlag(),
		},
		Action: download,
	}
	return cmd
}

func download(c *cli.Context) error {
	f := form.New(
		c.String(flags.REGION_FLAG_NAME),
	)
	err := f.Validate()
	if err != nil {
		return err
	}
	_log.Debug("download form values", zap.Object("form", f))
	return cmd.Apply(f)
}
