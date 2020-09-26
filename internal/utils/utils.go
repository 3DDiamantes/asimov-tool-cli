package utils

import (
	"os"
	"unicode"
)

func FileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
