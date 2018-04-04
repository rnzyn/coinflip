package core

import (
	"net/http"

	"github.com/labstack/echo"
)

type CoinflipContext struct {
	echo.Context
	Coinflip *Coinflip
}

func (ctx *CoinflipContext) JsonError(err error) error {
	return ctx.JSON(http.StatusInternalServerError, map[string]string{
		"error": err.Error(),
	})
}
