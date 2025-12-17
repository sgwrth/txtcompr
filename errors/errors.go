package errors

import (
	"log"
)

func MissingArg() {
	log.Fatal("Missing argument(s)")
}

func InvalidFilepath(err error) {
	log.Fatalf("Invalid filepath: %v", err)
}
