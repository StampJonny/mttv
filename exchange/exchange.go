package exchange

import (
	"github.com/stampjohnny/mttv/e"
	"github.com/stampjohnny/mttv/logging"
	"github.com/stampjohnny/mttv/utils"
)

type Exchange interface {
	SetCryptoBalance(float64) error
	GetCryptoBalance() float64
	SetMoneyBalance(float64) error
	GetMoneyBalance() float64
	SetPrice(float64) error
	GetPrice() float64
	Buy(float64) (interface{}, error)
	// GetName() string
}

var currentExchange Exchange

func SetTestExchange(callbacks ...interface{}) error {
	currentExchange = &TestExchange{}
	for _, callback := range callbacks {
		callbackFunc := callback.(func(interface{}))
		callbackFunc(currentExchange)
	}
	return nil
}

func SetExchange(stockName string) error {
	if exchange := GetExchange(exchangeName); exchange != nil {
		currentExchange = exchange
	} else {
		return e.Err("no exchange with name: %v", stockName)
	}
	return nil
}

func Buy(amount float64) (interface{}, error) {
	ctx, err := currentExchange.Buy(amount)
	if err != nil {
		return nil, e.Err("can't buy crypto: %v", err)
	}

	if err := logging.LogBuyContext(ctx.(interface {
		GetAmount() float64
		GetMoneyBalance() float64
	})); err != nil {
		utils.Crash("can't log operation %v", err)
	}

	return ctx, nil
}

func GetCryptoBalance() float64 {
	return currentExchange.GetCryptoBalance()
}

func SetCryptoBalance(amount float64) {
	currentExchange.SetCryptoBalance(amount)
}

func GetMoneyBalance() float64 {
	return currentExchange.GetMoneyBalance()
}

func SetMoneyBalance(amount float64) {
	currentExchange.SetMoneyBalance(amount)
}

func SetPrice(price float64) error {
	return currentExchange.SetPrice(price)
}

func GetPrice() float64 {
	return currentExchange.GetPrice()
}

// type TransactionID uuid.UUID

// func Buy() (uuid.UUID, error) {
// 	txn, _ := uuid.NewRandom()
// 	ctx := context.Get()
// 	amount := ctx.GetCryptoAmount()
// 	if err := currentExchange.Buy(amount); err != nil {
// 		return txn, fmt.Errorf("can't buy through exchange %s: %s", currentExchange.GetName(), err)

// 	}
// 	ctx.SetTransactionID(txn)
// 	return txn, nil
// }
