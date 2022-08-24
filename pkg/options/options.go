package options

import "github.com/urfave/cli/v2"

type Options struct {
	Host            string
	ApiKey          string
	ConfigDirectory string
	Insecure        bool
}

func GetOptions(c *cli.Context) *Options {
	return &Options{
		Host:            c.String("host"),
		ApiKey:          c.String("api-key"),
		ConfigDirectory: c.String("config-dir"),
		Insecure:        c.Bool("insecure"),
	}
}
