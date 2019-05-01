package context

import (
	"github.com/stampjohnny/mttv/i"
)

type contextImpl struct {
	i.Context

	Amount i.BalanceType
}

func (c *contextImpl) SetAmount(amount float64) error {
	c.Amount = i.BalanceType(amount)
	return nil
}
func (c *contextImpl) GetAmount() i.BalanceType {
	return c.Amount
}

func New() i.Context {
	return &contextImpl{}
}

var currentContext i.Context

func GetCurrentContext() i.Context {
	return currentContext
}

func SetCurrentContext(context i.Context) {
	currentContext = context
}
