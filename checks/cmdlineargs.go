package checks

import (
	"os"

	"dev.sgwrth/txtcompr/errors"
)

func IsArgPresent(args []string) {
	noArg := len(args) == 1
	argIsEmptyString := args[1] == ""
	if noArg || argIsEmptyString {
		errors.NoArgPresent()
	}
}

func FileExists(path string) {
	_, err := os.Stat(path)
	if err != nil {
		errors.InvalidFilepath(err)
	}
}
