package errors

import (
	"fmt"
)

func NoArgPresent() {
	fmt.Println("Missing path argument.")
}
