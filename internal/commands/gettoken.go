package commands

import (
	"asimov-tool-cli/internal/env"
	"fmt"

	"github.com/urfave/cli"
)

func GetToken(c *cli.Context) {
	token, err := env.GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Token: %s\n", token)
}
