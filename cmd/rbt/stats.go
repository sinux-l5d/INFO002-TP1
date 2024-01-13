package main

import (
	"fmt"
	"strconv"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
	"github.com/sinux-l5d/INFO002-TP1/internal/table"
	"github.com/urfave/cli/v2"
)

func init() {
	RegisterSubCmd(&cli.Command{
		Name:  "stats",
		Usage: "Print coverage of a rainbow table",
		Description: "Print the coverage of the table given the parameters or a filename.\n" +
			"If a filename is given, the global parameters are ignored.\n" +
			"If the width and height are given, global parameters are used.",
		ArgsUsage: "[<filename>|<width> <height>]",
		Action: func(c *cli.Context) error {
			// VALIDATE
			if c.Args().Len() == 1 {
				filename := c.Args().Get(0)
				t, err := table.Load(filename)
				if err != nil {
					return fmt.Errorf("cannot load table: %w", err)
				}
				fmt.Println(t.Config.String())
				fmt.Printf("Coverage: %.2f%%", t.Coverage())
			} else if c.Args().Len() == 2 {
				widthS := c.Args().Get(0)
				width, err := strconv.ParseUint(widthS, 10, 64)
				if err != nil {
					return fmt.Errorf("cannot parse width: %w", err)
				}

				heightS := c.Args().Get(1)
				height, err := strconv.ParseUint(heightS, 10, 64)
				if err != nil {
					return fmt.Errorf("cannot parse height: %w", err)
				}

				fmt.Println(config.GlobalConfig.String())
				fmt.Printf("Coverage: %.2f%%", table.Coverage(config.GlobalConfig, width, height))
			} else {
				cli.ShowSubcommandHelp(c)
				return fmt.Errorf("missing arguments")
			}
			return nil
		},
	})
}
