package main

import (
	"errors"
	"strconv"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
	"github.com/sinux-l5d/INFO002-TP1/internal/table"
	"github.com/urfave/cli/v2"
)

func init() {
	RegisterSubCmd(&cli.Command{
		Name:  "create",
		Usage: "Create a rainbow table",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "random",
				Aliases: []string{"r"},
				Usage:   "Randomize the index generated",
			},
		},
		ArgsUsage: "<width> <height> <filename>",
		Action: func(c *cli.Context) error {
			// VALIDATE

			random := c.Bool("random")
			largeurS := c.Args().Get(0)
			hauteurS := c.Args().Get(1)
			filename := c.Args().Get(2)

			if largeurS == "" || hauteurS == "" || filename == "" {
				cli.ShowSubcommandHelp(c)
				return errors.New("missing arguments")
			}

			largeur, err := strconv.ParseUint(largeurS, 10, 64)
			if err != nil {
				return errors.New("width not a valid uint64")
			}

			hauteur, err := strconv.ParseUint(hauteurS, 10, 64)
			if err != nil {
				return errors.New("height not a valid uint64")
			}

			// CREATE TABLE

			t, err := table.NewTable(config.GlobalConfig, largeur, hauteur, random)
			if err != nil {
				return err
			}

			return t.Save(filename)
		},
	})
}
