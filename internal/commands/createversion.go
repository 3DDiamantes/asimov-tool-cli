package commands

import (
	"fmt"
	"os"
	"regexp"

	"github.com/urfave/cli"
)

const (
	ProjectUnknown = iota
	ProjectGo
	ProjectNodeJS
	ProjectC
)

func CreateVersion(c *cli.Context) {
	if !c.Args().Present() {
		fmt.Println("Version number must be specified.")
		return
	}

	version := c.Args().First()

	if !isValidVersion(version) {
		fmt.Println("Incorrect version format.\nThe format must be mayor.minor.patch.\n")
		return
	}

	projectType := getProjectType()
	switch projectType {
	case ProjectGo:
		fmt.Println("Go project detected.")
	case ProjectUnknown:
		fmt.Println("Unknown project type.")
		return
	}

	fmt.Printf("Creating version %s\n", version)
}

func isValidVersion(v string) bool {
	re := regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)$`)

	results := re.FindString(v)

	if results != "" {
		return true
	}
	return false
}

func getProjectType() int {
	if _, err := os.Stat("go.mod"); err == nil {
		return ProjectGo
	}

	return ProjectUnknown
}
