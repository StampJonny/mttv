package main

import (
	"fmt"
	"time"

	"github.com/stampjohnny/mttv/config"
	"github.com/stampjohnny/mttv/trader"
)

func main() {
	checkSell := time.Tick(config.SellTimeoutSecond * time.Second)
	checkBuy := time.Tick(config.SellTimeoutSecond * time.Second)
	// getTransactions := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-checkSell:
			fmt.Printf("%+v\n", "selling...")
			trader.Sell()
		case <-checkBuy:
			trader.Buy()
			fmt.Printf("%+v\n", "buying")
		}
	}
}
