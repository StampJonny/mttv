package exchange

import (
	"github.com/stampjohnny/mttv/context"
	"github.com/stampjohnny/mttv/i"
)

var currentExchange i.Exchange

func SetExchange(exchange i.Exchange) error {
	currentExchange = exchange
	return nil
}

func Buy() error {
	context := context.GetCurrentContext()
	return currentExchange.Buy(context.GetAmount())
}

func GetBalance() i.Balance {
	return currentExchange.GetBalance()
}
