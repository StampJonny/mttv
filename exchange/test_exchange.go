package exchange

import (
	"errors"
	"sync"

	"github.com/stampjohnny/mttv/e"
	"github.com/stampjohnny/mttv/logging"
	"github.com/stampjohnny/mttv/utils"
	"golang.org/x/xerrors"
)

const exchangeName = "Test exchange"

type TestExchange struct {
	cryptoBalance     float64
	moneyBalance      float64
	cryptoAmountToBuy float64
	price             float64
	lock              sync.RWMutex
	freezed           bool
}

func (t *TestExchange) Copy() *TestExchange {
	newExchange := TestExchange{
		cryptoBalance:     t.GetCryptoBalance(),
		moneyBalance:      t.GetMoneyBalance(),
		cryptoAmountToBuy: t.GetAmount(),
		price:             t.GetPrice(),
		freezed:           true,
	}
	return &newExchange
}

func (t *TestExchange) isFreezed() bool {
	return t.freezed
}

func (t *TestExchange) GetAmount() float64 {
	logging.Debug("amount %v", t.cryptoAmountToBuy)
	return t.cryptoAmountToBuy
}
func (t *TestExchange) saveAmount(amount float64) {
	t.cryptoAmountToBuy = amount
}

func (t *TestExchange) Buy(amount float64) (interface{}, error) {
	logging.Debug("amount=%v", 1)
	if t.isFreezed() {
		return nil, e.Err("context is freezed")
	}
	t.lock.Lock()
	defer t.lock.Unlock()

	t.saveAmount(amount)

	if err := t.SetMoneyBalance(t.GetMoneyBalance() - t.getMoneyToPay()); err != nil {
		return nil, xerrors.Errorf("catn't update money balance: %v", err)
	}
	if err := t.SetCryptoBalance(t.GetCryptoBalance() + t.GetAmount()); err != nil {
		return nil, xerrors.Errorf("catn't update crypto balance: %v", err)
	}
	return t.Copy(), nil
}
func (t *TestExchange) getMoneyToPay() float64 {
	return t.cryptoAmountToBuy * t.price
}

func (t *TestExchange) SetDefault() {
	if t.isFreezed() {
		utils.Crash("context is freezed")
	}

	t.cryptoBalance = float64(100)
	t.moneyBalance = float64(100000)
	t.price = 5600
}

func (t *TestExchange) GetCryptoBalance() float64 {
	return t.cryptoBalance
}

func (t *TestExchange) SetCryptoBalance(amount float64) error {
	if t.isFreezed() {
		return e.Err("context is freezed")
	}

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
	if t.isFreezed() {
		return e.Err("context is freezed")
	}

	if amount <= 0 {
		return errors.New("can't set negative amount")
	}
	t.moneyBalance = amount
	return nil
}

func (t *TestExchange) SetPrice(price float64) error {
	if t.isFreezed() {
		return e.Err("context is freezed")
	}
	t.price = price
	return nil
}

func (t *TestExchange) GetPrice() float64 {
	return t.price
}

func (t *TestExchange) Save() error {
	return nil
}
