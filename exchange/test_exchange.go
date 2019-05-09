package exchange

import (
	"errors"
	"fmt"
)

const exchangeName = "Test exchange"

type TestExchange struct {
	cryptoBalance float64
	moneyBalance  float64
	price         float64
}

func (t *TestExchange) Buy(amount float64) error {
	moneyNeeded := amount * t.price
	err := t.SetMoneyBalance(t.GetMoneyBalance() - moneyNeeded)
	if err != nil {
		return fmt.Errorf("catn't update money balance: %s", err)
	}
	err = t.SetCryptoBalance(t.GetCryptoBalance() + amount)
	if err != nil {
		return fmt.Errorf("catn't update crypto balance: %s", err)
	}
	return nil
}

func (t *TestExchange) SetDefault() {
	t.cryptoBalance = float64(100)
	t.moneyBalance = float64(100000)
	t.price = 5600
}

func (t *TestExchange) GetCryptoBalance() float64 {
	return t.cryptoBalance
}

func (t *TestExchange) SetCryptoBalance(amount float64) error {
	if amount <= 0 {
		return errors.New("can't set negative amount")
	}
	t.cryptoBalance = amount
	return nil
}

func (t *TestExchange) GetMoneyBalance() float64 {
	return t.moneyBalance
}

func (t *TestExchange) SetMoneyBalance(amount float64) error {
	if amount <= 0 {
		return errors.New("can't set negative amount")
	}
	t.moneyBalance = amount
	return nil
}

func (t *TestExchange) SetPrice(price float64) {
	t.price = price
}

func (t *TestExchange) GetPrice() float64 {
	return t.price
}

func (t *TestExchange) Save() error {
	return nil
}
