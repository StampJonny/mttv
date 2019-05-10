package main

import (
	"fmt"
	"time"

	"github.com/stampjohnny/mttv/config"
	"github.com/stampjohnny/mttv/exchange"
	"github.com/stampjohnny/mttv/trader"
)

func main() {
	// checkSell := time.Tick(config.SellTimeoutSecond * time.Second)
	checkBuy := time.Tick(config.BuyTimeoutSecond * time.Second)
	// getTransactions := time.Tick(100 * time.Millisecond)
	err := exchange.SetExchange(config.ExchangeName)
	if err != nil {
		panic(fmt.Sprintf("Can't setup exchange %s; dyeing", err))
	}
	for {
		select {
		// case <-checkSell:
		// 	fmt.Printf("%+v\n", "selling...")
		// 	trader.Sell()
		case <-checkBuy:
			trader.Buy()
			fmt.Printf("%+v\n", "buying")
		}

	}
}
