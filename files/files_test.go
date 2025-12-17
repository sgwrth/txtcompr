package files

import (
	"testing"
)

func TestFilenameFromPath(t *testing.T) {
	filePath := "/home/me/docs/text.txt"
	filenameFromPath := FilenameFromPath(filePath)

	if filenameFromPath != "text" {
		t.Errorf("Filename is not 'text' but '%v'", filenameFromPath)
	}

	filePath = "/home/me/docs/text"
	filenameFromPath = FilenameFromPath(filePath)

	if filenameFromPath != "text" {
		t.Errorf("Filename is not 'text' but '%v'", filenameFromPath)
	}
}
