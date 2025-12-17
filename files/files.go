package files

import (
	"log"
	"os"
	"path/filepath"

	"dev.sgwrth/txtcompr/process"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}

func WriteCompressedFile(filePath string) {
	file := OpenFile(filePath)
	tokens := process.GetTokens(file)
	freqMap := process.FreqMap(tokens)
	duplicateEntries := process.DuplicateEntries(freqMap)
	dict := process.BuildDict(duplicateEntries)
	compressedTextBytes := process.CompressedTextBytes(tokens, dict)
	dictAsBytes := process.DictAsBinary(dict)
	WriteDictAndTextToFile(dictAsBytes, compressedTextBytes, "./output/arose2")
}

func FilenameFromPath(path string) string {
	filename := filepath.Base(path)
	var dotPos int
	dotFound := false

	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			dotFound = true
			dotPos = i
			break
		}
	}

	if dotFound {
		return filename[:dotPos]
	}

	return filename
}

func WriteDictAndTextToFile(dict []byte, compressedText []byte, filePath string) error {
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	_, err = file.Write(dict)

	if err != nil {
		return err
	}

	_, err = file.Write(compressedText)
	return err
}
