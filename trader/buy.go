package trader

import "github.com/stampjohnny/mttv/exchange"

func Buy() error {
	err := exchange.Buy()
	return err
}
