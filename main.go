package main

import (
	"coinbase-pro-trading-bot/common"
	"coinbase-pro-trading-bot/simplebot"
	"log"

	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/shopspring/decimal"
)

var coinId = "MANA"
var tradingPair = "MANA-USDC"

var buyPrice decimal.Decimal = decimal.NewFromFloat(1.07)
var sellPrice decimal.Decimal = decimal.NewFromFloat(1.14)

//for now, initialSize is determined by budget at buy price.
var initialSize decimal.Decimal = decimal.NewFromFloat(37)

func main() {

	var err error

	common.InitClient()
	common.InitCurrencies()

	//startOrder, err := common.Client.GetOrder("5357f73b-97fb-4668-9950-4de55eb7e064")

	//Start order is an initital order for the program to watch. This can be found using GetOrders and documenting the ID
	startOrder := coinbasepro.Order{ID: ""}

	err = simplebot.LimitOrdersConstantPrices(coinId, tradingPair, sellPrice, buyPrice, initialSize, startOrder, 3)

	log.Println(err.Error())
}
