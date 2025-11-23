package main

import (
	//	"fmt"
	"os"

	"dev.sgwrth/txtcompr/checks"
	// "dev.sgwrth/txtcompr/errors"
)

func main() {
	checks.IsArgPresent(os.Args)
	txtFilePath := os.Args[1]
	checks.FileExists(txtFilePath)
	// file = openTxtFile(path)
	// fileCompressed = compressFile(file)
	// saveFile(fileCompressed)
}
