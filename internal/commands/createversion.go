package commands

import (
	"asimov-tool-cli/internal/utils"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/urfave/cli"
)

const (
	ProjectUnknown = iota
	ProjectGo
)

func CreateVersion(c *cli.Context) {
	if !c.Args().Present() {
		fmt.Println("Version number must be specified.")
		return
	}

	version := c.Args().First()

	if !isValidVersion(version) {
		fmt.Println("Incorrect version format.\nThe format must be mayor.minor.patch.")
		return
	}

	target := c.String("target")
	if !isValidTarget(target) {
		fmt.Println("Target not valid.\nValid targets are: arm.")
		return
	}

	featureName, err := getFeatureName()
	if err != nil {
		fmt.Println("Error getting the feature name.")
		return
	}

	filename := version + "-" + featureName

	if utils.FileExist("builds/" + target + "/" + filename) {
		fmt.Println("The version already exists.\nPlease specify a different version number.")
		return
	}

	projectType := getProjectType()

	switch projectType {
	case ProjectGo:
		fmt.Println("Go project detected.")
		err = createGoVersion(filename, target)
	case ProjectUnknown:
		fmt.Println("Unknown project type.")
		return
	}

	if err != nil {
		fmt.Printf("Fail to create version.\nError: %v\n", err)
		return
	}

	fmt.Println("Version created successfully!")
}

func isValidVersion(v string) bool {
	re := regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)$`)

	results := re.FindString(v)

	if results != "" {
		return true
	}
	return false
}

func isValidTarget(target string) bool {
	if target == "current" || target == "arm" {
		return true
	}
	return false
}

func getProjectType() int {
	if utils.FileExist("go.mod") {
		return ProjectGo
	}

	return ProjectUnknown
}

func getFeatureName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	re := regexp.MustCompile(`feature/(.+)`)
	results := re.FindStringSubmatch(out.String())

	return results[1], nil
}

func createGoVersion(filename string, target string) error {
	fmt.Println("Creating version " + filename + " for " + target + "...")

	if target == "arm" {
		os.Setenv("GOOS", "linux")
		os.Setenv("GOARCH", "arm")
		os.Setenv("GOARM", "7")
	}

	cmd := exec.Command("go", "build", "-o", "builds/"+target+"/"+filename, "cmd/main.go")

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
