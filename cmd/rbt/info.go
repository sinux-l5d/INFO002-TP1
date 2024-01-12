package main

import (
	"errors"
	"fmt"

	"github.com/sinux-l5d/INFO002-TP1/internal/table"
	"github.com/urfave/cli/v2"
)

func init() {
	RegisterSubCmd(&cli.Command{
		Name:      "info",
		Usage:     "Print information about a rainbow table",
		ArgsUsage: "<filename>",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "Print all the content of the table",
			},
			&cli.IntFlag{
				Name:    "max",
				Aliases: []string{"m"},
				Usage:   "Maximum number of lines to print from the table",
				Value:   10,
			},
		},
		Action: func(c *cli.Context) error {
			filename := c.Args().Get(0)
			if filename == "" {
				cli.ShowSubcommandHelp(c)
				return errors.New("missing arguments")
			}

			t, err := table.Load(filename)
			if err != nil {
				return err
			}

			fmt.Printf("== Table %s ==\n", filename)
			fmt.Printf("Width: %d\n", t.Largeur)
			fmt.Printf("Height: %d\n", t.Hauteur)
			fmt.Printf("Alphabet: %s\n", t.Alphabet)
			fmt.Printf("Size: %d\n", t.Size)

			limit := c.Int("max")
			if c.Bool("all") {
				limit = 0
			}
			fmt.Printf("Content:\n%s\n", t.Print(limit))

			return nil
		},
	})
}
