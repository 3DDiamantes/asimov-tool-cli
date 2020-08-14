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
				Name:      "new-feature",
				Aliases:   []string{"nf"},
				Usage:     "Create a new feature branch",
				ArgsUsage: "featureName",
				Action:    commands.NewFeature,
			},
			{
				Name:      "create-version",
				Aliases:   []string{"cv"},
				Usage:     "Create a new version of the proyect.",
				ArgsUsage: "x.x.x",
				Action:    commands.CreateVersion,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "target",
						Value: "current",
						Usage: "OS and architecture for the build. If not set the current OS/Arch will be used.",
					},
				},
			},
			{
				Name:    "token-get",
				Aliases: []string{"tg"},
				Usage:   "Get your personal token for Github API",
				Action:  commands.GetToken,
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
