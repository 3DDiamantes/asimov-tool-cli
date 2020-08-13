package env

import (
	"errors"
	"os"
)

const tokenKey = "ASIMOV_TOOL_CLI_TOKEN"

func GetToken() (string, error) {
	token, exist := os.LookupEnv(tokenKey)
	if !exist || len(token) == 0 {
		return "", errors.New("Token not set.")
	}
	return token, nil
}
