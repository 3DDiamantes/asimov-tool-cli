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

	featureBranch = fmt.Sprintf("feature/%s", featureBranch)

	// git status -> must return nothing to commit
	fmt.Printf("Checking for pending changes to commit...")
	if pending, err := git.CommitsPending(); err != nil || pending {
		fmt.Printf("FAIL\nYou have pending changes to commit.\n")
		return
	}
	fmt.Printf("OK\n")

	// git checkout develop
	fmt.Printf("Checking out to develop...")
	if err := git.Checkout("develop"); err != nil {
		fmt.Printf("FAIL\n%v\n", err)
		return
	}
	fmt.Printf("OK\n")

	// git pull
	fmt.Printf("Downloading remote changes...")
	if err := git.Pull(); err != nil {
		fmt.Printf("FAIL\n%v\n", err)
		return
	}
	fmt.Printf("OK\n")

	// git checkout -b feature/featureName
	fmt.Printf("Creating the new branch and checking out...")
	if err := git.CreateBranch(featureBranch); err != nil {
		fmt.Printf("FAIL\n%v\n", err)
		return
	}
	fmt.Printf("OK\n")

	// git push -u origin feature/featureName
	fmt.Printf("Uploading the local branch...")
	if err := git.PushNewBranch(featureBranch); err != nil {
		fmt.Printf("FAIL\n%v\n", err)
		return
	}
	fmt.Printf("OK\n")

	// API call to Github
	fmt.Printf("Creating Pull Request...")
	if err := git.CreatePR("develop", featureBranch); err != nil {
		fmt.Printf("FAIL\n%v\n", err)
		return
	}
	fmt.Printf("OK\n")

	fmt.Printf("\nFeature '%s' created correctly. Happy coding!\n", featureBranch)
}
