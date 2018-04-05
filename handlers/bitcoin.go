package handlers

import (
	"errors"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/labstack/echo"
)

func BitcoinDonation(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)
	return ctx.JsonError(errors.New("Not implemented"))
}