package errors

import (
	"log"
)

func NoArgPresent() {
	log.Fatal("Error: missing filepath")
}

func InvalidFilepath(err error) {
	log.Fatalf("Error: %v", err)
}
