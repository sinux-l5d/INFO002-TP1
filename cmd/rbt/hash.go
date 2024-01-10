package main

import (
	"crypto/sha1"
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
)

func init() {
	RegisterSubCmd(&cli.Command{
		Name:      "hash",
		Usage:     "Demo hash command",
		UsageText: progname + " hash [command options] [string to hash]",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "algo",
				Aliases: []string{"a"},
				Usage:   "Hash algorithm. Options : sha1",
				Value:   "sha1",
			},
		},
		Action: func(c *cli.Context) error {
			if c.String("algo") != "sha1" {
				return errors.New("only sha1 is supported")
			}

			toHash := c.Args().First()
			if toHash == "" {
				cli.ShowSubcommandHelp(c)
				return errors.New("missing a string to hash")
			}

			h := sha1.New()
			h.Write([]byte(toHash))
			fmt.Printf("%X (%s)\n", h.Sum(nil), toHash)
			return nil
		},
	})
}
