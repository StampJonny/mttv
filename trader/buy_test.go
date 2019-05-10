package trader_test

import (
	"fmt"
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
				ctx := c.(interface {
					SetDefault()
					SetuptTestEnv()
				})
				ctx.SetDefault()
				ctx.SetuptTestEnv()
				s.ctx = c
			},
		),
	)

}

func (s *TestSuite) TestBuyCryptoIncreased() {
	cryptoAmountToBuy := 0.001
	prevCryptoBalance := exchange.GetCryptoBalance()

	_, err := exchange.Buy(cryptoAmountToBuy)
	s.NoError(err)

	expCrypto := prevCryptoBalance + cryptoAmountToBuy
	s.Equal(expCrypto, exchange.GetCryptoBalance())
}

func (s *TestSuite) TestBuyMoneyDecreased() {
	cryptoAmountToBuy := 0.001
	prevMoneyBalance := exchange.GetMoneyBalance()

	_, err := exchange.Buy(cryptoAmountToBuy)
	s.NoError(err)

	expMoney := prevMoneyBalance - cryptoAmountToBuy*exchange.GetPrice()
	s.Equal(expMoney, exchange.GetMoneyBalance())
}

func (s *TestSuite) TestBuyContextSaved() {
	logging.EnableDebug()
	cryptoAmountToBuy := 0.001

	log, err := logging.Get(logging.BuyLog)
	s.NoError(err)

	ctx, err := exchange.Buy(cryptoAmountToBuy)
	s.NoError(err)

	s.NoError(err)
	s.FileExists(log.GetFilepath())

	r, err := logging.Reader(log.GetFilepath())
	s.NoError(err)

	line, err := r.ReadLine()

	s.NoError(err)
	type i interface {
		GetAmount() float64
		GetMoneyBalance() float64
	}
	s.Contains(line,
		fmt.Sprintf(`"amount":"%v"`, ctx.(i).GetAmount()),
		log.GetFilepath())
	s.Contains(line,
		fmt.Sprintf(`"money":"%v"`, ctx.(i).GetMoneyBalance()),
		log.GetFilepath())

	line, err = r.ReadLine()
	s.Error(err)
}
