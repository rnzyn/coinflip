package handlers

import (
	"fmt"
	"net/http"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/ShoppersShop/coinflip/payloads"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"
)

func (h *Coinflip) WhitelistPost(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Parse payload
	payload := new(payloads.WhitelistAdd)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JsonError(err)
	}

	// Convert raw input to Ethereum addresses
	addresses := []common.Address{}
	for _, address := range payload.Addresses {
		if common.IsHexAddress(address) {
			addresses = append(addresses, common.HexToAddress(address))
		} else {
			err := fmt.Errorf(core.ErrNotEnoughConfirmations, address)
			return ctx.JsonError(err)
		}
	}

	// Send transaction
	transaction, err := h.Contract.WhitelistAdd(h.TxOpts, addresses)
	if err != nil {
		return ctx.JsonError(err)
	}

	// Return response
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"txHash": transaction.Hash().String(),
	})
}
