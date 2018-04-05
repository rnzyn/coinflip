package handlers

import (
	"net/http"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/labstack/echo"
)

func Healthcheck(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)
	return ctx.JSON(http.StatusOK, map[string]string{
		"status": "OK",
	})
}
