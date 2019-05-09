package logging

import (
	"bufio"
	"os"

	"github.com/stampjohnny/mttv/e"
)

type reader struct {
	file *os.File
	s    *bufio.Scanner
}

func Reader(filePath string) (*reader, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, e.Err("logger.Reader : %s", err)
	}

	r := reader{}
	if err := r.init(f); err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *reader) init(file *os.File) error {
	r.file = file
	r.s = bufio.NewScanner(file)
	return nil
}

func (r *reader) ReadLine() (string, error) {
	if r.s.Scan() {
		return r.s.Text(), nil
	}

	return "", e.Err("EOF")
}
