package main

import (
	"encoding/hex"
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
			h2i,
			i2i,
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
			&cli.BoolFlag{
				Name:    "lowercase",
				Aliases: []string{"l"},
				Usage:   "Lowercase the resulting hash",
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

			// PRINT
			r, err := test.Run()
			if err != nil {
				return err
			}

			if c.Bool("lowercase") {
				fmt.Printf("%x", r)
				return nil
			}

			fmt.Printf("%X", r)

			if config.GlobalConfig.Verbose {
				fmt.Printf(" (%s)\n", toHash)
			}
			return nil
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
		Name:        "i2c",
		Usage:       "Index to combination",
		Description: "Transform the input number into the corresponding combination, depending on the alphabet in global options",
		ArgsUsage:   "<index>",
		Action: func(c *cli.Context) error {
			// VALIDATE
			if c.Args().First() == "" {
				cli.ShowSubcommandHelp(c)
				return errors.New("missing the number")
			}

			// Convert string to int
			n, err := strconv.ParseUint(c.Args().First(), 10, 64)
			if err != nil {
				cli.ShowSubcommandHelp(c)
				return errors.New("invalid number")
			}

			// TEST
			test, err := tests.NewI2CTest(&config.GlobalConfig, n)
			if err != nil {
				return err
			}

			// PRINT
			r, err := test.Run()
			if err != nil {
				return err
			}
			fmt.Printf("i2c(%d)=%s\n", n, r)
			return nil
		},
	}

	h2i = &cli.Command{
		Name:        "h2i",
		Usage:       "Hash to index",
		Description: "Transform the input hexadecimal hash into a number for the given column (don't forget global options...)",
		ArgsUsage:   "<hash> <column>",
		Action: func(c *cli.Context) error {
			// VALIDATE
			if c.Args().First() == "" {
				cli.ShowSubcommandHelp(c)
				return errors.New("missing the string")
			}
			str := c.Args().First()

			if c.Args().Get(1) == "" {
				cli.ShowSubcommandHelp(c)
				return errors.New("missing the column number")
			}

			// Convert string to int
			i := c.Args().Get(1)
			t, err := strconv.ParseUint(i, 10, 64)
			if err != nil {
				cli.ShowSubcommandHelp(c)
				return errors.New("invalid column number")
			}

			// TEST
			hash, err := hex.DecodeString(str)
			if err != nil {
				return fmt.Errorf("invalid hash: %v", err)
			}

			test, err := tests.NewH2ITest(&config.GlobalConfig, hash, t)
			if err != nil {
				return err
			}

			// PRINT
			r, err := test.Run()
			if err != nil {
				return err
			}
			fmt.Printf("h2i(sha1(%s),%d)=%d\n", str, t, r)
			return nil
		},
	}

	i2i = &cli.Command{
		Name:        "i2i",
		Usage:       "Number to number",
		Description: "Convert a number to text that is hashed, then convert if back to another number for the given column (don't forget global options...)",
		ArgsUsage:   "<number> <column>",
		Action: func(c *cli.Context) error {
			// VALIDATE
			if c.Args().First() == "" {
				cli.ShowSubcommandHelp(c)
				return errors.New("missing the number")
			}

			// Convert string to int
			n, err := strconv.ParseUint(c.Args().First(), 10, 64)
			if err != nil {
				cli.ShowSubcommandHelp(c)
				return fmt.Errorf("invalid number: %v", err)
			}

			if c.Args().Get(1) == "" {
				cli.ShowSubcommandHelp(c)
				return errors.New("missing the column number")
			}

			// Convert string to int
			i := c.Args().Get(1)
			t, err := strconv.ParseUint(i, 10, 64)
			if err != nil {
				cli.ShowSubcommandHelp(c)
				return errors.New("invalid column number")
			}

			// TEST
			test, err := tests.NewI2ITest(&config.GlobalConfig, n, t)
			if err != nil {
				return err
			}

			// PRINT
			r, err := test.Run()
			if err != nil {
				return err
			}
			fmt.Printf("i2i(%d,%d)=%d\n", n, t, r)
			return nil
		},
	}
)
