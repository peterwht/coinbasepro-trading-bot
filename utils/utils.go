package utils

import (
	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/shopspring/decimal"
)

//returns amount of crypto after fees
func GetCryptoAmount(order coinbasepro.Order, coinId string) (decimal.Decimal, error) {

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
