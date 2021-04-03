package simplebot

import (
	"coinbase-pro-trading-bot/common"
	"coinbase-pro-trading-bot/utils"
	"errors"
	"time"

	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/shopspring/decimal"
)

/* LimitOrdersConstantPrices places limit orders at a set sell & buy price.
Saves last order, and checks if that order was filled. If a buy order is filled
then a sell (limit) order is placed, and vice versa.
*/
func LimitOrdersConstantPrices(coinId, tradingPair string, sellPrice, buyPrice, initialSize decimal.Decimal, startOrder coinbasepro.Order, delayMultiplier float32) error {

	var err error
	var currentOrder coinbasepro.Order = startOrder

	if currentOrder.ID == "" {
		currentOrder, err = utils.PlaceOrder(buyPrice.String(), initialSize.String(), "buy", tradingPair, true, currentOrder)

		if err != nil {
			return err
		}
	}

	for true {

		//get the currently placed order
		pendingOrder, err := common.Client.GetOrder(currentOrder.ID)
		if err != nil {
			return (err)
		}

		if pendingOrder.DoneReason == "filled" {
			//if it was filled, and was a buy order
			if currentOrder.Side == "buy" {
				//gets the amount of crypto coins after fees
				amount, err := utils.GetCryptoAmount(pendingOrder, coinId)
				if err != nil {
					return err
				}

				//place sell order with
				currentOrder, err = utils.PlaceOrder(sellPrice.String(), amount.String(), "sell", tradingPair, true, currentOrder)
				if err != nil {
					return err
				}
			} else { // if sell order
				amount, err := utils.GetCryptoAmount(pendingOrder, coinId)
				if err != nil {
					return err
				}
				currentOrder, err = utils.PlaceOrder(buyPrice.String(), amount.String(), "buy", tradingPair, true, currentOrder)
				if err != nil {
					return err
				}
			}
		} else if pendingOrder.ID == "" { // if the ID does not exist, the order no longer exists (and was not filled)
			return errors.New("Placed order no longer exits (probably cancelled")
		}

		time.Sleep(time.Second * time.Duration(delayMultiplier)) // delay. Not really necessary, but may be useful if running in go routine

	}

	return nil
}
