package utils

import (
	"fmt"
	"log"
)

func Assertf(expr bool, msg string, args ...any) {
	if !expr {
		log.Fatal(fmt.Sprintf(msg, args...))
	}
}
