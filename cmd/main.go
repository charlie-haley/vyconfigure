package cmd

import (
	"os"

	"github.com/charlie-haley/vyconfigure/pkg/options"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var appVersion = "development"

func Run() {
	o := options.Options{}

	app := cli.NewApp()
	app.Name = "vyconfigure"
	app.Version = appVersion
	app.Usage = "Declarative configuration for VyOS."
	app.EnableBashCompletion = true
	app.Authors = []*cli.Author{
		{Name: "Charlie Haley", Email: "charlie-haley@users.noreply.github.com"},
	}
	app.Flags = []cli.Flag{
		&cli.StringFlag{Destination: &o.Host, Name: "host", Usage: "The hostname of the VyOS HTTP API.", EnvVars: []string{"VYCONFIGURE_HOST"}},
		&cli.StringFlag{Destination: &o.ApiKey, Name: "api-key", Usage: "API key for the HTTP API.", EnvVars: []string{"VYCONFIGURE_API_KEY"}},
		&cli.StringFlag{Destination: &o.ConfigDirectory, Name: "config-dir", Value: ".", Usage: "Directory where config is stored.", EnvVars: []string{"VYCONFIGURE_CONFIG_DIR"}},
		&cli.BoolFlag{Destination: &o.Insecure, Name: "insecure", Usage: "Whether to skip verifying the SSL certificate.", EnvVars: []string{"VYCONFIGURE_INSECURE"}},
	}
	app.Commands = []*cli.Command{
		{
			Name: "version", Aliases: []string{"v"}, Usage: "prints the current version.",
			Action: version,
		},
		{
			Name: "apply", Aliases: []string{"a"}, Usage: "applies the current configuration.",
			Action: apply,
		},
		{
			Name: "sync", Aliases: []string{"s"}, Usage: "syncs configuration to the current directory through the HTTP API.",
			Action: sync,
		},
		{
			Name: "plan", Aliases: []string{"d"}, Usage: "shows the diff between the current directory and VyOS instance",
			Action: plan,
		},
	}
	app.Action = version

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
