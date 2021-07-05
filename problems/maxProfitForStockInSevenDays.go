package problems

import (
	model "probsolbygolang/model"
)

func FindMaxProfitForStockOverSevenDays(purchaseDates []model.StockPurchaseDates, stockPrice []int) model.StockPurchaseDates {
	var maxProfit int
	var dates model.StockPurchaseDates
	for i := range purchaseDates {
		profit := stockPrice[purchaseDates[i].Sell] - stockPrice[purchaseDates[i].Buy]
		if maxProfit < profit {
			dates = purchaseDates[i]
		}
	}
	return dates
}
