package commands

import (
	"asimov-tool-cli/internal/utils"
	"fmt"

	"strings"

	"asimov-tool-cli/internal/git"

	"github.com/urfave/cli"
)

const illegalCharacters string = " !\"#$%&'()*+,./:;<=>?@[\\]^_`{|}~\n\r"

func NewFeature(c *cli.Context) {
	// Validate arguments
	if !c.Args().Present() {
		fmt.Println("Feature's name must be specified.")
		return
	}

	featureBranch := c.Args().First()

	if strings.ContainsAny(featureBranch, illegalCharacters) {
		fmt.Println("Feature's name contains illegal characters.")
		return
	}

	featureBranch = fmt.Sprintf("feature/%s", featureBranch)

	// git status -> must return nothing to commit
	pending, err := git.CommitsPending()
	if err != nil {
		utils.PrintError("Failed to check for pending changes to commit", err)
		return
	}
	if pending {
		utils.PrintFail("You have pending changes to commit")
		return
	}

	utils.PrintOK("No pending changes to commit")

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

	fmt.Printf("\nFeature '%s' created correctly. Please create a PR in GitHub.\nHappy coding!\n", featureBranch)
}
