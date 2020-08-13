package env

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"
)

const tokenKey = "ASIMOV_TOOL_CLI_TOKEN"

func GetTokenCmd(c *cli.Context) {
	token, err := GetToken()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Token: %s\n", token)
}

func GetToken() (string, error) {
	token, exist := os.LookupEnv(tokenKey)
	if !exist || len(token) == 0 {
		return "", errors.New("Token not set.")
	}
	return token, nil
}
