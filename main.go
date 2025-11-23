package main

import (
	"fmt"
	"os"

	"dev.sgwrth/txtcompr/checks"
	"dev.sgwrth/txtcompr/errors"
)

func main() {
	pathArgPresent := checks.IsArgPresent(os.Args)
	if !pathArgPresent {
		errors.NoArgPresent()
		return
	}
	txtFilePath := os.Args[1]
	fmt.Println(txtFilePath)
	// file = openTxtFile(path)
	// fileCompressed = compressFile(file)
	// saveFile(fileCompressed)
}
