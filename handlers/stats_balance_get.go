package handlers

import (
	"fmt"
	"net/http"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/ShoppersShop/coinflip/responses"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"
)

func (h *Coinflip) BalanceGet(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Parse input
	address := common.HexToAddress(c.Param("address"))
	if !common.IsHexAddress(c.Param("address")) {
		err := fmt.Errorf(core.ErrInvalidEthereumAddress, address)
		return ctx.JsonError(err)
	}

	// Call smart contract
	balance, err := h.TokenContract.BalanceOf(nil, address)
	if err != nil {
		return ctx.JsonError(err)
	}

	return ctx.JSON(http.StatusOK, responses.Balance{
		Balance: balance.Uint64(),
	})
}
