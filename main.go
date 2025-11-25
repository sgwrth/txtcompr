package main

import (
	// "fmt"
	"os"

	"dev.sgwrth/txtcompr/checks"
	"dev.sgwrth/txtcompr/files"
	"dev.sgwrth/txtcompr/process"
)

func main() {
	checks.IsArgPresent(os.Args)
	txtFilePath := os.Args[1]
	checks.FileExists(txtFilePath)
	txtFile := files.OpenFile(txtFilePath)
	process.CompressFile(txtFile)
	// saveFile(fileCompressed)
}
