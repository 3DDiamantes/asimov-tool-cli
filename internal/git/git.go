package git

import (
	"asimov-tool-cli/internal/repository"
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
)

func CommitsPending() (bool, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return false, err
	}

	return out.String() != "", nil
}
func Checkout(branch string) error {
	cmd := exec.Command("git", "checkout", branch)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
func Pull() error {
	cmd := exec.Command("git", "pull")
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
func CreateBranch(branch string) error {
	cmd := exec.Command("git", "checkout", "-b", branch)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
func PushNewBranch(branch string) error {
	cmd := exec.Command("git", "push", "-u", "origin", branch)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
func CreatePR(baseBranch string, headBranch string) error {
	owner, repo, err := getOwnerAndRepository()
	if err != nil {
		return err
	}

	r := repository.NewGithubRepository(owner, repo)

	resp, err := r.CreatePR(
		baseBranch,
		headBranch,
		fmt.Sprintf("%s->%s", headBranch, baseBranch),
	)

	if err != nil {
		return err
	}
	if resp.StatusCode() != 201 {
		return errors.New("Pull request creation failed. Please create it manually.")
	}
	return nil
}
func getOwnerAndRepository() (string, string, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", "", err
	}

	re := regexp.MustCompile(`:(.+)/(.+)\.git`)
	results := re.FindStringSubmatch(out.String())

	return results[1], results[2], nil
}
func GetProject() (string, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", err
	}

	re := regexp.MustCompile(`:.+/(.+)\.git`)
	results := re.FindStringSubmatch(out.String())

	return results[0], nil
}
