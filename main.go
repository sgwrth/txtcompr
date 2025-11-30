package main

import (
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
	tokens := process.GetTokens(txtFile)
	freqMap := process.FreqMap(tokens)
	duplicateEntries := process.DuplicateEntries(freqMap)
	dict := process.BuildDict(duplicateEntries)
	compressedTextBytes := process.CompressedTextBytes(tokens, dict)
	files.SaveCompressedText(compressedTextBytes, "./output/arose")
	// process.CompressFile(txtFile)
	// saveFile(fileCompressed)
}
