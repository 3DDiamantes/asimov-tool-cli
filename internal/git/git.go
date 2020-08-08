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
