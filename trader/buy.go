package trader

import "github.com/stampjohnny/mttv/exchange"

func Buy() (interface{}, error) {
	return exchange.Buy(0.01)
}
