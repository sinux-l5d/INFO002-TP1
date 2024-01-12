package main

import (
	"errors"
	"fmt"

	"github.com/sinux-l5d/INFO002-TP1/internal/table"
	"github.com/urfave/cli/v2"
)

func init() {
	RegisterSubCmd(&cli.Command{
		Name:      "crack",
		Usage:     "Crack a hash using a rainbow table",
		ArgsUsage: "<hash> <filename>",
		Action: func(c *cli.Context) error {
			// VALIDATE
			v := c.Bool("verbose")
			hash := c.Args().Get(0)
			filename := c.Args().Get(1)

			if hash == "" || filename == "" {
				cli.ShowSubcommandHelp(c)
				return errors.New("missing arguments")
			}

			// CRACK

			t, err := table.Load(filename)
			if err != nil {
				return err
			}

			t.Config.Verbose = v // I know, I know, ugly but still for educational purposes
			candidat, err := t.Crack(hash)
			if err != nil {
				return err
			}

			fmt.Println(candidat)
			return nil
		},
	})
}
