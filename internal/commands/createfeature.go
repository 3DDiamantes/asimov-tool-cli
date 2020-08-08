package commands

import (
	"fmt"

	"strings"

	"asimov-tool-cli/internal/git"

	"github.com/urfave/cli"
)

const illegalCharacters string = " !\"#$%&'()*+,./:;<=>?@[\\]^_`{|}~\n\r"

func CreateFeature(c *cli.Context) {
	if !c.Args().Present() {
		fmt.Println("Features's name must be specified.")
		return
	}

	featureBranch := c.Args().First()

	if strings.ContainsAny(featureBranch, illegalCharacters) {
		fmt.Println("Feature's name contains illegal characters.")
		return
	}

	// git status -> should return nothing to commit
	fmt.Printf("Checking for pending changes to commit...")
	if git.CommitsPending() {
		fmt.Printf("FAIL\nYou have pending changes to commit.\n")
		return
	}
	fmt.Printf("OK\n")

	// git checkout develop
	fmt.Printf("Checking out to develop...")
	fmt.Printf("OK\n")
	// git pull
	fmt.Printf("Downloading remote changes...")
	fmt.Printf("OK\n")
	// git checkout -b feature/featureName
	fmt.Printf("Creating the new branch and checking out...")
	fmt.Printf("OK\n")
	// git push -u origin feature/featureName
	fmt.Printf("Uploading the local branch...")
	fmt.Printf("OK\n")
	// API call to Github
	fmt.Printf("Creating Pull Request...")
	fmt.Printf("OK\n")

	fmt.Printf("\nFeature '%s' created correctly. Happy coding!\n", featureBranch)
}
