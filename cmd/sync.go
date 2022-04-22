package cmd

import (
	"github.com/charlie-haley/vyconfigure/pkg/api"
	"github.com/charlie-haley/vyconfigure/pkg/config"
	"github.com/charlie-haley/vyconfigure/pkg/options"
	"github.com/urfave/cli/v2"
)

func sync(c *cli.Context) error {
	o := options.GetOptions(c)

	client, err := api.CreateClient(o)
	if err != nil {
		return err
	}

	d, err := client.Retrieve()
	if err != nil {
		return err
	}

	err = config.Write(d, o)
	if err != nil {
		return err
	}

	return nil
}
