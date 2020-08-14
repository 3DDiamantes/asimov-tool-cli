package commands

import (
	"asimov-tool-cli/internal/utils"
	"bytes"
	"fmt"
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

	featureName, err := getFeatureName()
	if err != nil {
		fmt.Println("Error getting the feature name.")
		return
	}

	filename := version + "-" + featureName
	fmt.Println("Creating version " + filename + "...")

	if utils.FileExist("builds/" + filename) {
		fmt.Println("The version already exists.\nPlease specify a different version number.")
		return
	}

	projectType := getProjectType()

	switch projectType {
	case ProjectGo:
		fmt.Println("Go project detected.")
		err = createGoVersion(filename)
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

func createGoVersion(filename string) error {
	//os.Setenv("GOOS", "linux")
	//os.Setenv("GOARCH", "arm")
	//os.Setenv("GOARM", "7")
	cmd := exec.Command("go", "build", "-o", "builds/"+filename, "cmd/main.go")

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
