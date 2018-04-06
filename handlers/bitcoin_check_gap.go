package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/ShoppersShop/coinflip/payloads"
	httpclient "github.com/ddliu/go-httpclient"
	"github.com/labstack/echo"
)

func (h *Coinflip) BitcoinCheckGap(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Perform API request
	requestUrl := core.BlockchainInfoBaseUrl + core.BlockchainInfoAddressGap
	res, err := httpclient.Get(requestUrl, map[string]string{
		"key":  h.Config.BlockchainInfoApiKey,
		"xpub": h.Config.BitcoinAccountXpub,
	})

	if err != nil {
		return ctx.JsonError(err)
	}

	// Read response body
	bodyBytes, err := res.ReadAll()
	if err != nil {
		return ctx.JsonError(err)
	}

	// Unmarshal response
	response := payloads.BlockchainGap{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return ctx.JsonError(err)
	}

	return ctx.JSON(http.StatusOK, response)
}
