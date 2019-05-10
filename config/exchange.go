package config

import "time"

const TestExchangeName = "TestExchange"

var ExchangeName = TestExchangeName
var BuyTimeoutSecond = time.Duration(10)

var BaseLogDir = "/mttv/"
