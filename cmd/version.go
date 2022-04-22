package cmd

import (
	"os"

	"github.com/urfave/cli/v2"
)

func version(c *cli.Context) error {
	println(appVersion)
	os.Exit(0)
	return nil
}
