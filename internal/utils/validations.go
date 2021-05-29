package utils

import (
	"asimov-tool-cli/internal/defines"
	"regexp"
	"unicode"
)

func IsLower(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) && !unicode.IsLower(r) {
			return false
		}
	}
	return true
}
func IsValidVersion(v string) bool {
	re := regexp.MustCompile(`^(\d+)\.(\d+)\.(\d+)$`)

	results := re.FindString(v)

	return results != ""
}
func IsValidTarget(target string) bool {
	return target == "" ||
		target == defines.TargetARM ||
		target == defines.TargetLinux ||
		target == defines.TargetMac ||
		target == defines.TargetWin
}
func IsValidVersionName(versionName string) bool {
	return IsLower(versionName)
}
