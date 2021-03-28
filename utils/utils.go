package utils

import (
	"github.com/preichenberger/go-coinbasepro/v2"
	"github.com/shopspring/decimal"
)

//coinbase pro is only accurate to 0.00000001
var decimalPlacesAllowed int32 = 8

//returns amount of crypto after fees
func GetCryptoAmount(order coinbasepro.Order) (string, error) {

	//total amount of fiat money used for order
	fiatGross, err := decimal.NewFromString(order.ExecutedValue)
	if err != nil {
		return "0", err
	}

	fees, err := decimal.NewFromString(order.FillFees)
	if err != nil {
		return "0", err
	}

	cryptoPrice, err := decimal.NewFromString(order.Price)

	fiatAfterFees := fiatGross.Sub(fees)

	//divide the amount of fiat by the price
	return fiatAfterFees.Div(cryptoPrice).Round(decimalPlacesAllowed).String(), nil
}
