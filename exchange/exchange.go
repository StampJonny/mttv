package exchange

import (
	"github.com/stampjohnny/mttv/e"
	"github.com/stampjohnny/mttv/logging"
)

type Exchange interface {
	SetCryptoBalance(float64) error
	GetCryptoBalance() float64
	SetMoneyBalance(float64) error
	GetMoneyBalance() float64
	SetPrice(float64)
	GetPrice() float64
	Buy(float64) error
	// GetName() string
}

var currentExchange Exchange

type Callback func(interface{})

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
		return e.Err("no exchange with name: %s", stockName)
	}
	return nil
}

func Buy(amount float64) error {
	ctx := currentExchange.Buy(amount)
	log, err := logging.Get(logging.TradingLog)
	if err != nil {
		return e.Err("can't get logger: %s", err)
	}

	log.Log(logging.Fields{"amount": amount}, "buy")

	return ctx
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

func SetPrice(price float64) {
	currentExchange.SetPrice(price)
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
