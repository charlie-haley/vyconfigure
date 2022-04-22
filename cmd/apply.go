package cmd

import (
	"github.com/charlie-haley/vyconfigure/pkg/api"
	"github.com/charlie-haley/vyconfigure/pkg/config"
	"github.com/charlie-haley/vyconfigure/pkg/convert"
	"github.com/charlie-haley/vyconfigure/pkg/options"
	r3diff "github.com/r3labs/diff/v3"
	"github.com/urfave/cli/v2"
)

func apply(c *cli.Context) error {
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
	changelog, err := r3diff.Diff(rc, lc)
	if err != nil {
		return err
	}

	var toDelete []string
	var toCreate []string
	if len(changelog) > 0 {
		for _, change := range changelog {
			if change.Type == "create" {
				toCreate = append(toCreate, change.To.(string))
			}
			if change.Type == "delete" {
				toDelete = append(toDelete, change.From.(string))
			}
		}
	} else {
		println("No changes to apply.")
		return nil
	}

	dc := convert.CmdsToData(toDelete, "delete")
	cc := convert.CmdsToData(toCreate, "set")

	cmds := append(dc, cc...)
	err = client.Configure(cmds)
	if err != nil {
		return err
	}

	return nil
}
