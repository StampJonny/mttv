package trader_test

import (
	"testing"

	"github.com/stampjohnny/mttv/exchange"
	"github.com/stampjohnny/mttv/logging"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite

	ctx interface{}
}

func TestRun(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (s *TestSuite) SetupTest() {
	s.NoError(
		exchange.SetTestExchange(
			func(c interface{}) {
				ctx := c.(interface{ SetDefault() })
				ctx.SetDefault()
				s.ctx = c
			},
		),
	)

}

func (s *TestSuite) TestBuyCryptoIncreased() {
	cryptoAmountToBuy := 0.001
	prevCryptoBalance := exchange.GetCryptoBalance()

	s.NoError(exchange.Buy(cryptoAmountToBuy))

	expCrypto := prevCryptoBalance + cryptoAmountToBuy
	s.Equal(expCrypto, exchange.GetCryptoBalance())
}

func (s *TestSuite) TestBuyMoneyDecreased() {
	cryptoAmountToBuy := 0.001
	prevMoneyBalance := exchange.GetMoneyBalance()

	s.NoError(exchange.Buy(cryptoAmountToBuy))

	expMoney := prevMoneyBalance - cryptoAmountToBuy*exchange.GetPrice()
	s.Equal(expMoney, exchange.GetMoneyBalance())
}

func (s *TestSuite) TestBuyContextSaved() {
	log, err := logging.Get(logging.TradingLog)
	defer func() { s.NoError(log.Remove()) }()

	s.NoError(exchange.Buy(0.001))

	s.NoError(err)
	s.FileExists(log.GetFilepath())

	r, err := logging.Reader(log.GetFilepath())
	s.NoError(err)

	line, err := r.ReadLine()

	s.NoError(err)
	s.Contains(line, `amount":0.001`)

	line, err = r.ReadLine()
	s.Error(err)
}
