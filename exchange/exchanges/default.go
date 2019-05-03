package exchanges

const exchangeName = "Test exchange"

type TestExchange struct {
	MoneyBalance  *testBalance
	CryptoBalance *testBalance
}

type testBalance struct {
	Amount float64
}

func newTestBalance(amount float64) *testBalance {
	return &testBalance{
		Amount: float64(amount),
	}
}

func (tb *testBalance) Available() float64 {
	return tb.Amount
}

func (tb *testBalance) Update(amount float64) {
	tb.Amount = tb.Amount + amount
}

func NewTest() *TestExchange {
	return &TestExchange{
		MoneyBalance:  newTestBalance(2),
		CryptoBalance: newTestBalance(2),
	}
}

func (t *TestExchange) GetMoneyBalance() float64 {
	return t.MoneyBalance.Available()
}

func (t *TestExchange) GetCryptoBalance() float64 {
	return t.CryptoBalance.Available()
}

func (t *TestExchange) GetName() string {
	return "Test Exchange"
}

func (t *TestExchange) Buy(amountCrypto float64) error {
	t.CryptoBalance.Update(amountCrypto)
	amountMoney := amountCrypto
	t.MoneyBalance.Update(-amountMoney)
	return nil
}
