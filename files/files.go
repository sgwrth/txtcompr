package files

import (
	"log"
	"os"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	return file
}
