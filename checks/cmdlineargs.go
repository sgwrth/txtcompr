package checks

import (
	"os"
	"slices"

	"dev.sgwrth/txtcompr/constants"
	"dev.sgwrth/txtcompr/errors"
)

func AreArgsPresent(args []string) {
	missingArg := len(args) < 2
	argIsEmptyString := args[1] == "" || args[2] == ""
	if missingArg || argIsEmptyString {
		errors.MissingArg()
	}
}

func FileExists(path string) {
	_, err := os.Stat(path)
	if err != nil {
		errors.InvalidFilepath(err)
	}
}

func IsValidOperationArg(arg string) bool {
	return slices.Contains(constants.AllOperationArgs(), arg)
}
