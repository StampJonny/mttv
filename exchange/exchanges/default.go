package exchanges

import (
	"github.com/stampjohnny/mttv/i"
)

type test struct {
	i.Exchange

	MoneyBalance  i.Balance
	CryptoBalance i.Balance
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

func NewTest() i.Exchange {
	return &test{
		MoneyBalance:  newTestBalance(2),
		CryptoBalance: newTestBalance(2),
	}
}

func (t *test) GetMoneyBalance() i.Balance {
	return t.MoneyBalance
}
func (t *test) GetCryptoBalance() i.Balance {
	return t.CryptoBalance
}

func (t *test) Buy(amountCrypto i.BalanceType) error {
	t.CryptoBalance.Update(amountCrypto)
	amountMoney := amountCrypto
	t.MoneyBalance.Update(-amountMoney)
	return nil
}
