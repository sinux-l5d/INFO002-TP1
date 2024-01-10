package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	app = &cli.App{
		Name:  "rbt",
		Usage: "Programme to manage a rainbow table",
	}
	progname = "rbt"
)

func RegisterSubCmd(cmd *cli.Command) {
	app.Commands = append(app.Commands, cmd)
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
