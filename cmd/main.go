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
				Name:      "build",
				Usage:     "Create a build of the project",
				ArgsUsage: "",
				Action:    commands.Build,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "target",
						Value: "",
						Usage: "OS and architecture for the build. If not set the current OS/Arch will be used. Supported targets: arm, linux, mac, win.",
					},
					&cli.StringFlag{
						Name:  "name",
						Value: "",
						Usage: "Name of the build. If not set the name will be project-branch-M.m.p",
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
