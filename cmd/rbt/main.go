package main

import (
	"log"
	"os"

	"github.com/sinux-l5d/INFO002-TP1/internal/config"
	"github.com/urfave/cli/v2"
)

const (
	progname = "rbt"
)

var (
	cfg = config.Config{}
	app = &cli.App{
		Name:  "rbt",
		Usage: "Programme to manage a rainbow table",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "alphabet",
				Usage:       "Alphabet to use",
				Value:       "abcdefghijklmnopqrstuvwxyz",
				Destination: &cfg.CustomAlphabet,
			},
			&cli.StringFlag{
				Name:        "abc",
				Aliases:     []string{"A"},
				Usage:       "Select a predefined alphabet. Possible values:\n" + config.Alphabets(),
				Destination: &cfg.Abc,
			},
			&cli.IntFlag{
				Name:        "size",
				Aliases:     []string{"s"},
				Usage:       "Size of the strings to generate",
				Value:       4,
				Destination: &cfg.Size,
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Usage:       "Verbose mode",
				Destination: &cfg.Verbose,
			},
		},
	}
)

func RegisterSubCmd(cmd *cli.Command) {
	app.Commands = append(app.Commands, cmd)
}

func main() {
	cfg.Writer = app.Writer
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
