package main

import (
	"log"
	"os"
	"sort"

	"asimov-tool-cli/internal/commands"
	"asimov-tool-cli/internal/env"

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
				Name:    "token-get",
				Aliases: []string{"tg"},
				Usage:   "Get your personal token for Github API",
				Action:  env.GetTokenCmd,
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
