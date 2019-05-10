package utils

import "flag"

func IsTestingMode() bool {
	if flag.Lookup("test.v") != nil {
		return true
	}
	return false
}
