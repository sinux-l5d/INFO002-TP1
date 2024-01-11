package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
	"github.com/sinux-l5d/INFO002-TP1/internal/tests"
	"github.com/urfave/cli/v2"
)

func init() {
	RegisterSubCmd(&cli.Command{
		Name:      "test",
		Usage:     "Run a test",
		UsageText: progname + " [global options] test <test name>",
		Subcommands: []*cli.Command{
			hash,
			i2c,
			globconfig},
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
			test, err := tests.NewHashTest(&config.GlobalConfig, c.String("algo"), toHash)
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
			fmt.Println(config.GlobalConfig.String())
			return nil
		},
	}

	i2c = &cli.Command{
		Name:      "i2c",
		Usage:     "Number to combination",
		ArgsUsage: "<number>",
		Action: func(c *cli.Context) error {
			// VALIDATE
			if c.Args().First() == "" {
				cli.ShowSubcommandHelp(c)
				return errors.New("missing the number")
			}

			// Convert string to int
			n, err := strconv.Atoi(c.Args().First())
			if err != nil {
				cli.ShowSubcommandHelp(c)
				return errors.New("invalid number")
			}

			// TEST
			test, err := tests.NewI2CTest(&config.GlobalConfig, n)
			if err != nil {
				return err
			}

			return test.Run()
		},
	}
)
