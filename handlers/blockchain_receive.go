package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/ShoppersShop/coinflip/models"
	"github.com/ShoppersShop/coinflip/payloads"
	"github.com/ShoppersShop/coinflip/responses"
	httpclient "github.com/ddliu/go-httpclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"
	uuid "github.com/satori/go.uuid"
)

func (h *Coinflip) BlockchainReceive(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Find available account
	account := models.Account{}
	if h.Database.Where("gap < ?", core.Bip44AddressLimit).First(&account).RecordNotFound() {
		return ctx.JsonError(errors.New(core.ErrNoAvailableAccountsFound))
	}

	// Parse payload
	payload := new(payloads.ReceiveAddress)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JsonError(err)
	}

	// Validate beneficiary address
	if !common.IsHexAddress(payload.Beneficiary) {
		return ctx.JsonError(fmt.Errorf(core.ErrInvalidEthereumAddress, payload.Beneficiary))
	}

	// Check if beneficiary already has pending transfer
	transfer := models.Transfer{}
	if !h.Database.Preload("Address").Where("beneficiary = ?", payload.Beneficiary).First(&transfer).RecordNotFound() {
		if transfer.Address == nil {
			return ctx.JsonError(errors.New(core.ErrUnknown))
		}
		return ctx.JSON(http.StatusOK, transfer.Address)
	}

	// Generate invoice ID
	invoiceID := uuid.NewV4()

	// Send API request
	requestUrl := core.BlockchainInfoBaseUrl + core.BlockchainInfoReceive
	res, err := httpclient.Get(requestUrl, map[string]string{
		"key":      h.Config.BlockchainInfoApiKey,
		"xpub":     account.Xpub,
		"callback": core.GetCallbackUrl(h.Config.Domain, invoiceID.String()),
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
	response := responses.BlockchainInfoReceive{}
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		return ctx.JsonError(err)
	}

	// Check for errors
	if response.Message != nil {
		return ctx.JsonError(errors.New(*response.Message))
	}

	// Begin transaction
	tx := h.Database.Begin()

	// Increase gap
	account.Gap = account.Gap + 1
	if err := h.Database.Save(&account).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Create new address
	address := models.Address{
		Address:   response.Address,
		AccountID: account.ID,
	}

	// Save new address
	if err := h.Database.Create(&address).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Create new transfer
	transfer = models.Transfer{
		InvoiceID:   invoiceID.String(),
		Beneficiary: payload.Beneficiary,
		AddressID:   address.ID,
	}

	// Save new transfer
	if err := h.Database.Create(&transfer).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Commit transaction
	tx.Commit()
	address.Transfers = append(address.Transfers, transfer)
	return ctx.JSON(http.StatusOK, address)
}
