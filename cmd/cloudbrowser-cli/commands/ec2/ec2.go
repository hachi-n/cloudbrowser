package ec2

import (
	"github.com/hachi-n/cloudbrowser/cmd/cloudbrowser-cli/commands/ec2/download"
	"github.com/urfave/cli/v2"
)

func Cmd() *cli.Command {
	cmd := &cli.Command{
		Name:  "ec2",
		Usage: "ec2 commands",
		Subcommands: []*cli.Command{
			download.Cmd(),
		},
	}
	return cmd
}
