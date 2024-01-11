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
	app = &cli.App{
		Name:  "rbt",
		Usage: "Programme to manage a rainbow table",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "alphabet",
				Usage:       "Alphabet to use",
				Value:       "abcdefghijklmnopqrstuvwxyz",
				Destination: &config.GlobalConfig.CustomAlphabet,
			},
			&cli.StringFlag{
				Name:        "abc",
				Aliases:     []string{"A"},
				Usage:       "Select a predefined alphabet. Possible values:\n" + config.Alphabets(),
				Destination: &config.GlobalConfig.Abc,
			},
			&cli.IntFlag{
				Name:        "size",
				Aliases:     []string{"s"},
				Usage:       "Size of the strings to generate",
				Value:       4,
				Destination: &config.GlobalConfig.Size,
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Usage:       "Verbose mode",
				Destination: &config.GlobalConfig.Verbose,
			},
		},
	}
)

func RegisterSubCmd(cmd *cli.Command) {
	app.Commands = append(app.Commands, cmd)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
