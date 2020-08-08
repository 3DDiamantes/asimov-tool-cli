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
	fmt.Println("Checking for pending changes to commit")
	if git.CommitsPending() {
		fmt.Println("You have pending changes to commit.")
		return
	}
	// git checkout develop
	fmt.Println("Checking out to develop")
	// git pull
	fmt.Println("Downloading remote changes")
	// git checkout -b feature/featureName
	fmt.Println("Creating the new branch and checking out")
	// git push -u origin feature/featureName
	fmt.Println("Uploading the local branch")
	// API call to Github
	fmt.Println("Creating Pull Request")

	fmt.Printf("Feature %s created correctly. Happy coding!\n", featureBranch)
}
