package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

func CommitsPending() bool {
	cmd := exec.Command("git", "status", "--porcelain")
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return true
	}

	return out.String() != ""
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
