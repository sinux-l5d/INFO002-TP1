package main

import (
	"errors"

	"github.com/urfave/cli/v2"
)

func init() {
	RegisterSubCmd(&cli.Command{
		Name: "create",
		Action: func(c *cli.Context) error {
			return errors.New("not implemented yet")
		},
	})
}
