package handlers

import (
	"net/http"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/labstack/echo"
)

func (h *Coinflip) StatsGet(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Retrieve data from smart contracts
	active, err := h.Contract.IsActiveSale(nil)
	availableUnits, err := h.Contract.AvailableUnits(nil)
	basePrice, err := h.Contract.BasePrice(nil)
	discountPrice, err := h.Contract.DiscountPrice(nil)
	duration, err := h.Contract.Duration(nil)
	minPayment, err := h.Contract.MinPayment(nil)
	startTime, err := h.Contract.StartTime(nil)
	unitsSold, err := h.Contract.UnitsSold(nil)
	weiReceived, err := h.Contract.WeiReceived(nil)

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
