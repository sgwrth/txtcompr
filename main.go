package main

import (
	"os"

	"dev.sgwrth/txtcompr/checks"
	"dev.sgwrth/txtcompr/files"
	"dev.sgwrth/txtcompr/process"
)

func main() {
	checks.AreArgsPresent(os.Args)
	filePath := os.Args[1]
	checks.FileExists(filePath)

	files.WriteCompressedFile(filePath)
	process.DecompressFile("./output/arose2")
}
