package utils

import (
	"coinbase-pro-trading-bot/common"

	"github.com/preichenberger/go-coinbasepro/v2"
)

func PlaceOrder(price, size, side, tradingPair string, postOnly bool, currentOrder coinbasepro.Order) (coinbasepro.Order, error) {
	orderConfig := coinbasepro.Order{
		Price:     price,
		Size:      size,
		Side:      side, // side is buy or sell
		ProductID: tradingPair,
		PostOnly:  postOnly,
	}

	tempOrder, err := common.Client.CreateOrder(&orderConfig)
	if err != nil {
		return currentOrder, err
	}

	return tempOrder, nil
}
