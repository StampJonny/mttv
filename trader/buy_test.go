package trader_test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
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

var availableMoneyBalance float64
var availableCryptoBalance float64

func (s *TestSuite) TestBuyPositive() {
	s.mockExchange()
	s.prepareContext()
	s.saveBalance()

	s.NoError(context.Init())

	txid := s.givenTraderBoughtCrypto()
	s.NoError(context.Save())

	s.accountMoneyBalanceDecreased()
	s.accountCryptoBalanceIncreased()
	//cryptoAmountSavedWithSellPrice()
	s.NoError(context.Find(txid.String()))
	ctx := context.Get()
	s.Equal(txid.String(), ctx.GetTransactionID().String())
	s.Equal(0.001, ctx.GetCryptoAmount())
	// s.Equal(ctx.GetAmount(), tx.GetCryptoAmount())
	// s.NotNil(tx.GetPrice())
	// s.NotNil(tx.GetMoneyAmount())
}

func (s *TestSuite) accountCryptoBalanceIncreased() {
	exp := availableCryptoBalance + 0.001
	s.Equal(exp, exchange.GetCryptoBalance())
}

func (s *TestSuite) accountMoneyBalanceDecreased() {
	exp := availableMoneyBalance - 0.001
	s.Equal(exp, exchange.GetMoneyBalance())
}

func (s *TestSuite) givenTraderBoughtCrypto() uuid.UUID {
	txID, err := trader.Buy()
	s.NoError(err)
	s.NotNil(txID)
	return txID
}

func (s *TestSuite) prepareContext() {
	ctx := context.New()
	s.NotNil(ctx)
	ctx.SetCryptoAmount(0.001)

	context.Set(ctx)
}

func (s *TestSuite) mockExchange() {
	stock := exchanges.NewTest()
	fmt.Printf("%+v\n", stock)                    // output for debug
	fmt.Printf("%+v\n", stock.GetName())          // output for debug
	fmt.Printf("%+v\n", stock.GetMoneyBalance())  // boutput for debug
	fmt.Printf("%+v\n", stock.GetCryptoBalance()) // output for debug

	s.NoError(exchange.SetExchange(stock))
}

func (s *TestSuite) saveBalance() {
	availableMoneyBalance = exchange.GetMoneyBalance()
	availableCryptoBalance = exchange.GetCryptoBalance()
}
