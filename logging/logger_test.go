package logging_test

import (
	"testing"

	"github.com/stampjohnny/mttv/logging"
	"github.com/stretchr/testify/suite"
)

type test struct {
	suite.Suite
}

type someContext struct{}

func (s *someContext) ToJson() string {
	return `{"some":"json"}`
}

func TestLogger(t *testing.T) {
	suite.Run(t, new(test))
}

func (s *test) TestLogRecord() {
	l, err := logging.Get("some-logger")
	s.NoError(err)
	l.Log(logging.Fields{"test": "field"}, "test")
}
