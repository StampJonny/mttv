package exchanges

import (
	"github.com/stampjohnny/mttv/i"
)

type test struct {
	i.Exchange

	Balance i.Balance
}

func NewTest() i.Exchange {
	return &test{
		Balance: newTestBalance(1),
	}
}

type testBalance struct {
	i.Balance

	Amount float64
}

func newTestBalance(amount float64) i.Balance {
	return &testBalance{
		Amount: amount,
	}
}

func (tb *testBalance) Available() float64 {
	return tb.Amount
}

func (tb *testBalance) Update(amount float64) {
	tb.Amount = tb.Amount + amount
}

func (t *test) GetBalance() i.Balance {
	return t.Balance
}

func (t *test) Buy(amount float64) error {
	t.Balance.Update(-amount)
	return nil
}
