package utils

import "fmt"

func Crash(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}
