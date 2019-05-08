package exchange

import (
	"github.com/stampjohnny/mttv/config"
)

func GetExchange(exchangeName string) Exchange {
	switch exchangeName {
	case config.TestExchangeName:
		return &TestExchange{}
	}
	return nil
}
