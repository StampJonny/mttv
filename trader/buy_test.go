package trader_test

import (
	"testing"

	"github.com/stampjohnny/mttv/context"
	"github.com/stampjohnny/mttv/exchange"
	"github.com/stampjohnny/mttv/exchange/exchanges"
	"github.com/stampjohnny/mttv/i"
	"github.com/stampjohnny/mttv/trader"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestRun(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

var availableBalance i.BalanceType

func (s *TestSuite) TestBuyPositive() {
	s.mockExchange()
	s.prepareContext()
	s.saveBalance()

	s.givenTraderBoughtCrypto()

	s.accountMoneyBalanceDecreased()
}

func (s *TestSuite) accountMoneyBalanceDecreased() {
	exp := availableBalance - 0.001
	s.Equal(exp, exchange.GetBalance().Available())
}

func (s *TestSuite) givenTraderBoughtCrypto() {
	s.NoError(trader.Buy())
}

func (s *TestSuite) prepareContext() {
	ctx := context.New()
	s.NotNil(ctx)
	s.NoError(ctx.SetAmount(0.001))

	context.SetCurrentContext(ctx)
}

func (s *TestSuite) mockExchange() {
	s.NoError(exchange.SetExchange(exchanges.NewTest()))
}

func (s *TestSuite) saveBalance() {
	availableBalance = exchange.GetBalance().Available()
}
