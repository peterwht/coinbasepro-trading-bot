package utils

import (
	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/shopspring/decimal"
)

func GetCryptoAmount(price, budget decimal.Decimal, coinId string) decimal.Decimal {
	return budget.Div(price).Truncate(GetCoinPrecision(coinId))
}

//returns amount of crypto after fees
func GetCryptoAmountAfterFees(order coinbasepro.Order, coinId string) (decimal.Decimal, error) {

	//total amount of fiat money used for order
	fiatGross, err := decimal.NewFromString(order.ExecutedValue)
	if err != nil {
		return decimal.NewFromInt32(0), err
	}

	fees, err := decimal.NewFromString(order.FillFees)
	if err != nil {
		return decimal.NewFromInt32(0), err
	}

	cryptoPrice, err := decimal.NewFromString(order.Price)

	fiatAfterFees := fiatGross.Sub(fees)

	//Each coin has a maximum accuracy for orders.
	//GetCoinPrecision gives the amount of decimal places
	coinPrecision := GetCoinPrecision(coinId)

	//divide the amount of fiat by the price
	return fiatAfterFees.Div(cryptoPrice).Truncate(coinPrecision), nil
}

func GetPercentIncrease(val1, val2 decimal.Decimal) decimal.Decimal {
	return (val2.Sub(val1)).Div(val1)
}

func FindClosestPrice(price decimal.Decimal, prices []decimal.Decimal) int {
	var minDiffIndex int = 0
	minDiff := (prices[minDiffIndex].Sub(price)).Abs()

	var diff decimal.Decimal
	for i, v := range prices[1:] {

		diff = (v.Sub(price)).Abs()
		if diff.LessThan(minDiff) {
			minDiff = diff
			minDiffIndex = i + 1 //+1 because we start at index 1
		}
	}

	return minDiffIndex
}
