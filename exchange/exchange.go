package exchange

import "fmt"

type Exchange interface {
	SetCryptoBalance(float64) error
	GetCryptoBalance() float64
	SetMoneyBalance(float64) error
	GetMoneyBalance() float64
	SetPrice(float64)
	Buy(float64) error
	// GetName() string
}

var currentExchange Exchange

type Callback func(Exchange)

func SetTestExchange(callbacks ...interface{}) error {
	currentExchange = &TestExchange{}
	for _, callback := range callbacks {
		callbackFunc := callback.(func(Exchange))
		callbackFunc(currentExchange)
	}
	return nil
}

func SetExchange(stockName string) error {
	if exchange := GetExchange(exchangeName); exchange != nil {
		currentExchange = exchange
	} else {
		return fmt.Errorf("no exchange with name: %s", stockName)
	}
	return nil
}

func Buy(amount float64) error {
	return currentExchange.Buy(amount)
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

// type Available interface {
// 	Available() float64
// }

// func GetMoneyBalance() float64 {
// 	return currentExchange.GetMoneyBalance()
// }
// func GetCryptoBalance() float64 {
// 	return currentExchange.GetCryptoBalance()
// }
