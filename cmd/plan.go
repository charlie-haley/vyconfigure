package cmd

import (
	"github.com/charlie-haley/vyconfigure/pkg/api"
	"github.com/charlie-haley/vyconfigure/pkg/config"
	"github.com/charlie-haley/vyconfigure/pkg/convert"
	"github.com/charlie-haley/vyconfigure/pkg/options"
	"github.com/fatih/color"
	diff "github.com/r3labs/diff/v3"
	"github.com/urfave/cli/v2"
)

func plan(c *cli.Context) error {
	o := options.GetOptions(c)

	// get remote config as cmds
	client, err := api.CreateClient(o)
	if err != nil {
		return err
	}

	d, err := client.RetrieveJson()
	if err != nil {
		return err
	}

	rc, _ := convert.JsonToCmds(d, "")

	// get local config as cmds
	lc, err := config.ReadAsCmds(o)
	if err != nil {
		return err
	}

	// get diff
	changelog, err := diff.Diff(rc, lc)
	if err != nil {
		return err
	}

	if len(changelog) > 0 {
		println("Changes to be applied:")
		for _, change := range changelog {
			if change.Type == "create" {
				color.Green("+ set " + change.To.(string))
			}
			if change.Type == "delete" {
				color.Red("- delete " + change.From.(string))
			}
		}
	} else {
		println("No changes to apply.")
	}

	return nil
}
