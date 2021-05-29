package commands

import (
	"asimov-tool-cli/internal/defines"
	"asimov-tool-cli/internal/git"
	"asimov-tool-cli/internal/utils"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/urfave/cli"
)

func Build(c *cli.Context) {
	// Check target
	target := c.String("target")
	if !utils.IsValidTarget(target) {
		utils.PrintFail("Invalid target. Valid targets are: arm, linux, mac, win.")
		return
	}
	if target == "" {
		target = utils.GetCurrentTarget()
	}
	utils.PrintOK("Target selected: " + target)

	// Check build name
	buildname := c.String("name")

	if !utils.IsValidVersionName(buildname) {
		utils.PrintFail("Invalid build name. Valid name should be lowercase letters only.")
		return
	}

	if buildname == "" {
		buildname, _ = generateBuildName()
	}

	utils.PrintOK("Build name: " + buildname)
	/*
		if filename == "" {
			featureName, err := getFeatureName()
			if err != nil {
				fmt.Println(err)
				return
			}

			filename = featureName + "-" + version
		}

		if utils.FileExist("builds/" + target + "/" + filename) {
			fmt.Println("The version already exists.\nPlease specify a different version number.")
			return
		}

		var err error
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

		fmt.Println("Version created successfully!")*/
}

// Utils
func getFeatureName() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	re := regexp.MustCompile(`feature/(.+)`)
	results := re.FindStringSubmatch(out.String())
	if len(results) < 1 {
		return "", errors.New("Builds can only be created from feature's branch")
	}

	return results[1], nil
}
func createGoVersion(filename string, target string) error {
	fmt.Println("Creating version " + filename + " for " + target + "...")

	switch target {
	case defines.TargetARM:
		os.Setenv(defines.EnvsGOOS, defines.OSARM)
		os.Setenv(defines.EnvsGOARCH, defines.ArchARM)
		os.Setenv(defines.EnvsGOARM, defines.ARMVersion)
	case defines.TargetLinux:
		os.Setenv(defines.EnvsGOOS, defines.OSLinux)
		os.Setenv(defines.EnvsGOARCH, defines.ArchLinux)
	case defines.TargetMac:
		os.Setenv(defines.EnvsGOOS, defines.OSMac)
		os.Setenv(defines.EnvsGOARCH, defines.ArchMac)
	case defines.TargetWin:
		os.Setenv(defines.EnvsGOOS, defines.OSWin)
		os.Setenv(defines.EnvsGOARCH, defines.ArchWin)
	}

	cmd := exec.Command("go", "build", "-o", "builds/"+target+"/"+filename, "cmd/main.go")

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func generateBuildName() (string, error) {
	project, err := git.GetProject()
	if err != nil {
		return "", err
	}
	branch := ""
	version := ""
	filename := fmt.Sprintf("%s_%s-v%s", project, branch, version)

	// Check if version already exist

	// If exist, increase version patch number

	return filename, nil
}
