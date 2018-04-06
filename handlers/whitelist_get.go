package handlers

import (
	"net/http"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"
)

func (h *Coinflip) WhitelistGet(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	result, err := h.Contract.WhitelistCheck(nil, common.HexToAddress(c.Param("address")))
	if err != nil {
		return ctx.JsonError(err)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"whitelisted": result,
	})
}
