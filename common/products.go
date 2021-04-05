package common

import (
	"github.com/preichenberger/go-coinbasepro/v2"
)

var Products map[string]coinbasepro.Product = make(map[string]coinbasepro.Product)

func InitProducts() error {
	products, err := Client.GetProducts()
	if err != nil {
		return err
	}

	for _, product := range products {
		Products[product.ID] = product
	}

	return nil
}
