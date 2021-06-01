package utils

import (
	"asimov-tool-cli/internal/defines"
	"fmt"
	"os"
	"runtime"
)

func FileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func PrintColor(text string, color string) string {
	colorStr := ""
	switch color {
	case "red":
		colorStr = "\033[31m"
	case "green":
		colorStr = "\033[32m"
	case "yellow":
		colorStr = "\033[33m"
	case "blue":
		colorStr = "\033[34m"
	}
	return fmt.Sprintf("%s%s%s", colorStr, text, "\033[0m")
}
func PrintError(message string, err error) {
	fmt.Printf("[%s] %s\nError: %s\n", PrintColor("ERROR", "red"), message, err.Error())
}
func PrintFail(message string) {
	fmt.Printf("[%s] %s\n", PrintColor("FAIL", "red"), message)
}
func PrintOK(message string) {
	fmt.Printf("[%s] %s\n", PrintColor("OK", "green"), message)
}

func GetCurrentTarget() string {
	arch := runtime.GOARCH
	os := runtime.GOOS

	if os == defines.OSWin && arch == defines.ArchWin {
		return defines.TargetWin
	}
	if os == defines.OSMac && arch == defines.ArchMac {
		return defines.TargetMac
	}
	if os == defines.OSLinux && arch == defines.ArchLinux {
		return defines.TargetLinux
	}
	if os == defines.OSARM && arch == defines.ArchARM {
		return defines.TargetARM
	}

	return ""
}

func GetProjectType() int {
	if FileExist("go.mod") {
		return defines.ProjectTypeGo
	}

	return defines.ProjectTypeUnknown
}
