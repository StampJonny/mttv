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
		Balance: newTestBalance(2),
	}
}

type testBalance struct {
	i.Balance

	Amount i.BalanceType
}

func newTestBalance(amount float64) i.Balance {
	return &testBalance{
		Amount: i.BalanceType(amount),
	}
}

func (tb *testBalance) Available() i.BalanceType {
	return tb.Amount
}

func (tb *testBalance) Update(amount i.BalanceType) {
	tb.Amount = tb.Amount + amount
}

func (t *test) GetBalance() i.Balance {
	return t.Balance
}

func (t *test) Buy(amount i.BalanceType) error {
	t.Balance.Update(-amount)
	return nil
}
