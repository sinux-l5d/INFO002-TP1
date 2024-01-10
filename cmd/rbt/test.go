package main

import (
	"errors"
	"fmt"

	"github.com/sinux-l5d/INFO002-TP1/internal/tests"
	"github.com/urfave/cli/v2"
)

func init() {
	RegisterSubCmd(&cli.Command{
		Name:        "test",
		Usage:       "Run a test",
		UsageText:   progname + " [global options] test <test name>",
		Subcommands: []*cli.Command{hash, globconfig},
	})
}

var (
	hash = &cli.Command{

		Name:  "hash",
		Usage: "Hash a string",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "algo",
				Aliases: []string{"a"},
				Usage:   "Hash algorithm. Options : sha1",
				Value:   "sha1",
			},
		},
		ArgsUsage: "<string to hash>",
		Action: func(c *cli.Context) error {
			// VALIDATE
			if c.String("algo") != "sha1" {
				return errors.New("only sha1 is supported")
			}

			toHash := c.Args().First()
			if toHash == "" {
				cli.ShowSubcommandHelp(c)
				return errors.New("missing a string to hash")
			}

			// TEST
			test, err := tests.NewHashTest(cfg, c.String("algo"), toHash)
			if err != nil {
				return err
			}

			return test.Run()
		},
	}

	globconfig = &cli.Command{
		Name:  "config",
		Usage: "Print the current configuration",
		Action: func(c *cli.Context) error {
			fmt.Println(cfg.String())
			return nil
		},
	}
)
