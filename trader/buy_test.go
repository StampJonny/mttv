package trader_test

import (
	"testing"

	"github.com/stampjohnny/mttv/context"
	"github.com/stampjohnny/mttv/exchange"
	"github.com/stampjohnny/mttv/exchange/exchanges"
	"github.com/stampjohnny/mttv/trader"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestRun(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) TestBuyPositive() {
	ctx := context.New()
	s.NotNil(ctx)
	s.NoError(ctx.SetAmount(0.001))

	context.SetCurrentContext(ctx)

	s.NoError(exchange.SetExchange(exchanges.NewTest()))
	availble := exchange.GetBalance().Available()

	s.NoError(trader.Buy())

	exp := availble - 0.001
	s.Equal(exp, exchange.GetBalance().Available())
}
