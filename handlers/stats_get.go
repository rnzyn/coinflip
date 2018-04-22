package handlers

import (
	"net/http"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/labstack/echo"
)

func (h *Coinflip) StatsGet(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Retrieve data from smart contracts
	active, err := h.SaleContract.IsActiveSale(nil)
	availableUnits, err := h.SaleContract.AvailableUnits(nil)
	basePrice, err := h.SaleContract.BasePrice(nil)
	discountPrice, err := h.SaleContract.DiscountPrice(nil)
	duration, err := h.SaleContract.Duration(nil)
	minPayment, err := h.SaleContract.MinPayment(nil)
	proxyAddress, err := h.SaleContract.ProxyAddress(nil)
	startTime, err := h.SaleContract.StartTime(nil)
	unitsSold, err := h.SaleContract.UnitsSold(nil)
	walletAddress, err := h.SaleContract.WalletAddress(nil)
	weiReceived, err := h.SaleContract.WeiReceived(nil)

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
		"proxyAddress":   proxyAddress,
		"startTime":      startTime,
		"unitsSold":      unitsSold,
		"walletAddress":  walletAddress,
		"weiReceived":    weiReceived,
	})
}
