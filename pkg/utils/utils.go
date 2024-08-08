package utils

import "fmt"

func Assertf(expr bool, msg string, args ...any) {
	if !expr {
		panic(fmt.Sprintf(msg, args...))
	}
}
