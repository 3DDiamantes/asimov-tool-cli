package main

import (
	"log"
	"os"
	"sort"

	"asimov-tool-cli/internal/commands"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Commands: []cli.Command{
			{
				Name:    "create-version",
				Aliases: []string{"cv"},
				Usage:   "Creates a version",
				Action:  commands.CreateVersion,
			},
			{
				Name:      "new-feature",
				Aliases:   []string{"nf"},
				Usage:     "Create a new feature branch",
				ArgsUsage: "featureName",
				Action:    commands.CreateFeature,
			},
		},
		Name:  "asimov",
		Usage: "CLI tool for manage builds, tests, deploys and other stuff related to development.",
	}

	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
