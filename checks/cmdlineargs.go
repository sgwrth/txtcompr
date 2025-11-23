package checks

import (
	"os"
)

func IsArgPresent(args []string) bool {
	if len(args) < 2 {
		return false
	}
	return true
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return true
	}
	return false
}
