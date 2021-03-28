package main

import (
	"coinbase-pro-trading-bot/common"
	"coinbase-pro-trading-bot/simplebot"

	"github.com/shopspring/decimal"
)

var tradingPair = "BTC-USD"

//example prices for sandbox. These prices do not make gains
var buyPrice decimal.Decimal = decimal.NewFromFloat(56002)
var sellPrice decimal.Decimal = decimal.NewFromFloat(56040)

//for now, 1.78 BTC was determined by BTC at 56002 using $100,000. 100,000/56002 = 1.78
var initialSize decimal.Decimal = decimal.NewFromFloat(1.78)

func main() {

	common.InitClient()
	err := simplebot.LimitOrdersConstantPrices(tradingPair, sellPrice, buyPrice, initialSize, 3)

	println(err.Error())

}
