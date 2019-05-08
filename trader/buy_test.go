package trader_test

import (
	"fmt"
	"testing"

	"github.com/stampjohnny/mttv/exchange"
	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func TestRun(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func cryptoBalance(amount float64) {
	fmt.Printf("%+v\n", amount) // output for debug

}
func (s *TestSuite) TestBuyCryptoIncreased() {
	currentCryptoBalance := float64(1)
	s.NoError(
		exchange.SetTestExchange(
			func(c exchange.Exchange) {
				s.NoError(c.SetMoneyBalance(10))
				c.SetPrice(1)
				s.NoError(c.SetCryptoBalance(currentCryptoBalance))
			},
		),
	)
	cryptoAmountToBuy := 0.001

	s.NoError(exchange.Buy(cryptoAmountToBuy))

	expCrypto := currentCryptoBalance + cryptoAmountToBuy
	s.Equal(expCrypto, exchange.GetCryptoBalance())
}

func (s *TestSuite) TestBuyMoneyDecreased() {
	currentPrice := float64(5600)
	currentMoneyBalance := float64(10000)
	s.NoError(
		exchange.SetTestExchange(
			func(c exchange.Exchange) {
				s.NoError(c.SetMoneyBalance(currentMoneyBalance))
				c.SetPrice(currentPrice)
			},
		),
	)
	cryptoAmountToBuy := 0.001

	s.NoError(exchange.Buy(cryptoAmountToBuy))

	expMoney := currentMoneyBalance - cryptoAmountToBuy*currentPrice
	s.Equal(expMoney, exchange.GetMoneyBalance())
}
