package checks

import (
	"testing"

	"dev.sgwrth/txtcompr/constants"
)

func TestOperationArg(t *testing.T) {
	if !IsValidOperationArg(constants.CompressArg()) {
		t.Error("Argument '-c' not recognized as valid")
	}

	if !IsValidOperationArg(constants.DecompressArg()) {
		t.Error("Argument '-d' not recognized as valid")
	}

	if IsValidOperationArg("-e") {
		t.Error("Argument '-e' recognized as valid")
	}

	if IsValidOperationArg("") {
		t.Error("Argument '' recognized as valid")
	}

	if IsValidOperationArg(" ") {
		t.Error("Argument ' ' recognized as valid")
	}

	if IsValidOperationArg("123") {
		t.Error("Argument '123' recognized as valid")
	}
}
