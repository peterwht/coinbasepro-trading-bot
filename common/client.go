package common

import "github.com/preichenberger/go-coinbasepro/v2"

var Client *coinbasepro.Client

func InitClient() {
	Client = coinbasepro.NewClient()

	/*
		BaseURL:
			https://api-public.sandbox.pro.coinbase.com
			or
			https://api.pro.coinbase.com

	*/
	Client.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL:    "<API Url>",
		Key:        "<Key>",
		Passphrase: "<Passphrase>",
		Secret:     "<Secret>",
	})
}
