package exchange

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/stampjohnny/mttv/context"
)

type Exchange interface {
	GetMoneyBalance() float64
	GetCryptoBalance() float64
	Buy(float64) error
	GetName() string
}

var currentExchange Exchange

func SetExchange(stock Exchange) error {
	currentExchange = stock
	return nil
}

type TransactionID uuid.UUID

func Buy() (uuid.UUID, error) {
	txn, _ := uuid.NewRandom()
	ctx := context.Get()
	amount := ctx.GetCryptoAmount()
	if err := currentExchange.Buy(amount); err != nil {
		return txn, fmt.Errorf("can't buy through exchange %s: %s", currentExchange.GetName(), err)

	}
	ctx.SetTransactionID(txn)
	return txn, nil
}

type Available interface {
	Available() float64
}

func GetMoneyBalance() float64 {
	return currentExchange.GetMoneyBalance()
}
func GetCryptoBalance() float64 {
	return currentExchange.GetCryptoBalance()
}
