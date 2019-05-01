package i

type Market int

type Account interface {
	SetExchange(Exchange) error
}

type Context interface {
	GetMarket() Market
	SetAmount(float64) error
	GetAmount() float64
}

type Balance interface {
	Available() float64
	Update(float64)
}

type Exchange interface {
	GetBalance() Balance
	Buy(amount float64) error
}
