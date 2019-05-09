package e

import "golang.org/x/xerrors"

func Err(format string, args ...interface{}) error {
	return xerrors.Errorf(format, args)
}
