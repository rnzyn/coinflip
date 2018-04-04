package handlers

import (
	"net/http"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/labstack/echo"
)

func GetStats(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Retrieve data from smart contracts
	active, err := ctx.Coinflip.Sale().IsActiveSale(nil)
	availableUnits, err := ctx.Coinflip.Sale().AvailableUnits(nil)
	basePrice, err := ctx.Coinflip.Sale().BasePrice(nil)
	discountPrice, err := ctx.Coinflip.Sale().DiscountPrice(nil)
	duration, err := ctx.Coinflip.Sale().Duration(nil)
	minPayment, err := ctx.Coinflip.Sale().MinPayment(nil)
	startTime, err := ctx.Coinflip.Sale().StartTime(nil)
	unitsSold, err := ctx.Coinflip.Sale().UnitsSold(nil)
	weiReceived, err := ctx.Coinflip.Sale().WeiReceived(nil)

	if err != nil {
		return ctx.JsonError(err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"active":         active,
		"availableUnits": availableUnits,
		"basePrice":      basePrice,
		"discountPrice":  discountPrice,
		"duration":       duration,
		"minPayment":     minPayment,
		"startTime":      startTime,
		"unitsSold":      unitsSold,
		"weiReceived":    weiReceived,
	})
}
