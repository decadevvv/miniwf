package utils

import (
	"fmt"
)

func PanicOnError(message string, err error) {
	if err != nil {
		panic(fmt.Errorf("%s: %w", message, err))
	}
}
