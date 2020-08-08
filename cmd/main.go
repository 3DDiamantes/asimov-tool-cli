package main

import (
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Commands: []cli.Command{
			{
				Name:    "create-version",
				Aliases: []string{"cv"},
				Usage:   "Creates a version",
				Action:  commands.createVersion,
			},
		},
		Name:  "asimov",
		Usage: "CLI tool for manage builds, tests, deploys and other stuff related to development.",
		Action: func(c *cli.Context) error {
			fmt.Println(c.Args())
			return nil
		},
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
