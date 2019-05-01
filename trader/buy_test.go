package trader_test

import (
	"testing"

	"github.com/stampjohnny/mttv/context"
	"github.com/stampjohnny/mttv/signal"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestRun(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestBuy() {
	ctx := context.New()

	signal, err := signal.Find(ctx)
	s.shouldBeValidSignal(err, signal)
	s.shoudlBeBuySignal(signal)

}

func (s *TestSuite) shoudlBeBuySignal(signal signal.Signal) {
	s.True(signal.IsBuySignal())
}

func (s *TestSuite) shouldBeValidSignal(err error, signal signal.Signal) {
	s.NoError(err)
	s.NotNil(signal)
}
