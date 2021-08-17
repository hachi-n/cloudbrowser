package flags

import (
	"github.com/urfave/cli/v2"
)

const (
	REGION_FLAG_NAME = "region"
)

func RegionFlag() *cli.StringFlag {
	return &cli.StringFlag{
		Name:     REGION_FLAG_NAME,
		Value:    "",
		Usage:    "region name.",
		Required: false,
	}
}
