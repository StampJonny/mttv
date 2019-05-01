package i

type Market int

type Account interface {
	SetExchange(Exchange) error
}

type Context interface {
	GetMarket() Market
	SetAmount(float64) error
	GetAmount() BalanceType
}

type BalanceType float64

type Balance interface {
	Available() BalanceType
	Update(BalanceType)
}

type Exchange interface {
	GetMoneyBalance() Balance
	GetCryptoBalance() Balance
	Buy(amount BalanceType) error
}
