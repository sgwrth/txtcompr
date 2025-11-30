package files

import (
	"log"
	"os"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	return file
}

func SaveCompressedText(bytes []byte, path string) {
	var permissions os.FileMode = 0644
	err := os.WriteFile(path, bytes, permissions)
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
}
