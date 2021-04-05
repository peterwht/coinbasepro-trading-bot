package gridbot

import (
	"coinbase-pro-trading-bot/common"
	"coinbase-pro-trading-bot/utils"
	"errors"
	"fmt"
	"time"

	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/shopspring/decimal"
)

func StartGridBot(coinId, tradingPair string, budget, lower, upper, gridLines decimal.Decimal, delayMultiplier int) error {
	span := upper.Sub(lower)
	priceChange := span.Div(gridLines.Sub(decimal.NewFromInt32(1)))
	orderAmount := budget.Div(gridLines)

	linePrices := []decimal.Decimal{lower}

	for i := decimal.NewFromInt32(1); i.LessThan(gridLines); i = i.Add(decimal.NewFromInt32(1)) {
		linePrices = append(linePrices, lower.Add(i.Mul(priceChange)).Round(utils.GetPricePrecision(tradingPair)))
	}

	fmt.Println("\nPrice/order: ", orderAmount)
	fmt.Println("Price/grid: ", priceChange)
	fmt.Println(linePrices)

	book, err := common.Client.GetBook(tradingPair, 1)
	if err != nil {
		return err
	}

	lastPrice, err := decimal.NewFromString(book.Bids[0].Price)
	if err != nil {
		return err
	}

	fmt.Println("Last Price: ", lastPrice)

	minDiffIndex := utils.FindClosestPrice(lastPrice, linePrices)
	// if the index is greater than the half point, then set it to the half point
	if minDiffIndex > int(len(linePrices)/2) {
		minDiffIndex = int(len(linePrices) / 2)
	}

	fmt.Println(linePrices[minDiffIndex])
	//time.Sleep(time.Second * 120)
	currentOrder, err := utils.PlaceOrder(
		linePrices[minDiffIndex].String(),
		utils.GetCryptoAmount(linePrices[minDiffIndex], orderAmount, coinId).String(),
		"buy",
		tradingPair,
		true,
		coinbasepro.Order{ID: ""})
	if err != nil {
		return err
	}

	currentOrders := []coinbasepro.Order{currentOrder}

	for true {

		for i, order := range currentOrders {
			//get the current order
			pendingOrder, err := common.Client.GetOrder(order.ID)
			if err != nil {
				return err
			}

			if pendingOrder.DoneReason == "filled" {
				lowerPrice, err := decimal.NewFromString(pendingOrder.Price)
				if err != nil {
					return err
				}
				upperPrice, err := decimal.NewFromString(pendingOrder.Price)
				if err != nil {
					return err
				}

				lowerPrice = lowerPrice.Sub(priceChange).Round(utils.GetPricePrecision(tradingPair))
				upperPrice = upperPrice.Add(priceChange).Round(utils.GetPricePrecision(tradingPair))

				fmt.Println(lowerPrice, " ", upperPrice)

				if !lowerPrice.LessThan(lower) {
					currentOrder, err = utils.PlaceOrder(lowerPrice.String(), utils.GetCryptoAmount(lowerPrice, orderAmount, coinId).String(), "buy", tradingPair, true, coinbasepro.Order{})
					if err != nil {
						return err
					}

					currentOrders = append(currentOrders, currentOrder)
				}

				if !upperPrice.GreaterThan(upper) {
					currentOrder, err = utils.PlaceOrder(upperPrice.String(), utils.GetCryptoAmount(upperPrice, orderAmount, coinId).String(), "sell", tradingPair, true, coinbasepro.Order{})
					if err != nil {
						return err
					}

					currentOrders = append(currentOrders, currentOrder)
				}

				//set one of the newly appended orders as the order we just checked.
				//range currentOrders is a copy, so it is safe to edit the slice

				currentOrders[i] = currentOrders[len(currentOrders)-1]
				//remove the order that we moved. : operator is exclusive on the upper bound
				currentOrders = currentOrders[0 : len(currentOrders)-1]

			} else if pendingOrder.ID == "" { // if the ID does not exist, the order no longer exists (and was not filled)
				return errors.New("Placed order no longer exists (probably cancelled")
			}

			time.Sleep(time.Second * time.Duration(delayMultiplier)) // delay
		}

	}

	return nil
}
