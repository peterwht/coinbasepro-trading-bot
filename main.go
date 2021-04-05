package main

import (
	"coinbase-pro-trading-bot/common"
	"coinbase-pro-trading-bot/gridbot"
	"fmt"
	"log"

	"github.com/shopspring/decimal"
)

var coinId = "MANA"
var tradingPair = "MANA-USDC"

var buyPrice decimal.Decimal = decimal.NewFromFloat(0.92)
var sellPrice decimal.Decimal = decimal.NewFromFloat(1.1)
var budget decimal.Decimal = decimal.NewFromFloat(100)

//for now, initialSize is determined by budget at buy price.
var initialSize decimal.Decimal = decimal.NewFromFloat(37)

func main() {

	var err error

	common.InitClient()

	err = common.InitCurrencies()
	if err != nil {
		log.Println(err)
	}
	err = common.InitAccountIDs()
	if err != nil {
		log.Println(err)
	}
	err = common.InitProducts()
	if err != nil {
		log.Println(err)
	}

	err = gridbot.StartGridBot(coinId, tradingPair, budget, buyPrice, sellPrice, decimal.NewFromInt32(8), 4)
	fmt.Println(err)

}
