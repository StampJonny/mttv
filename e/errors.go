package e

import (
	"golang.org/x/xerrors"
)

func Err(format string, args ...interface{}) error {
	err := xerrors.Errorf(format, args)
	return err
}
