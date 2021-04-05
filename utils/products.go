package utils

import "coinbase-pro-trading-bot/common"

//Returns the amount of decimal places allowed for a specific coin order
func GetPricePrecision(tradingPair string) int32 {
	minSizeStr := common.Products[tradingPair].QuoteIncrement

	var i int32 = 0
	for _, num := range minSizeStr {
		if num == '1' {
			break
		} else if num == '.' {
			continue
		}

		i++
	}

	return i
}
