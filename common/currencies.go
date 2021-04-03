package common

import "github.com/preichenberger/go-coinbasepro/v2"

var Currencies map[string]coinbasepro.Currency

//Gets information on all of the currencies.
//Returns a map with the coin id as the key and coinbasepro.Currency as the value
func InitCurrencies() error {
	Currencies = make(map[string]coinbasepro.Currency)

	currencies, err := Client.GetCurrencies()

	if err != nil {
		return err
	}

	for _, currency := range currencies {
		Currencies[currency.ID] = currency
	}
	return nil
}
