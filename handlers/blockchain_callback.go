package handlers

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/ShoppersShop/coinflip/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

func (h *Coinflip) BlockchainCallback(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)

	// Parse number of confirmations
	confirmations, err := strconv.Atoi(c.QueryParam("confirmations"))
	if err != nil {
		return ctx.JsonError(err)
	}

	// Validate required confirmations
	if confirmations < core.Confirmations {
		message := fmt.Errorf(core.ErrNotEnoughConfirmations, core.Confirmations, confirmations)
		return ctx.JsonError(message)
	}

	// Parse transfer value
	value, err := strconv.ParseInt(c.QueryParam("value"), 10, 64)
	if err != nil {
		return ctx.JsonError(err)
	}

	// Find transfer by invoice
	invoiceID := c.QueryParam("invoice_id")
	transfer := models.Transfer{}
	if h.Database.Preload("Address.Account").Where("invoice_id = ?", invoiceID).First(&transfer).RecordNotFound() {
		return ctx.JsonError(errors.New(core.ErrTransferNotFound))
	}

	// Calculate value in wei
	bitcoinValue := float64(value) / core.OneBitcoinInSatoshi
	etherValue := float64(bitcoinValue) * h.Config.BtcEthFallbackRate
	weiValue := etherValue * core.OneEtherInWei
	floatValue := big.NewFloat(weiValue).Text('g', 20)
	transferValue, success := new(big.Int).SetString(floatValue, 10)
	if !success {
		return ctx.JsonError(fmt.Errorf(core.ErrBtcEthConversionFailure, value))
	}

	// Debug calculations
	log.WithFields(log.Fields{
		"satoshi":  value,
		"bitcoin":  bitcoinValue,
		"ether":    etherValue,
		"wei":      weiValue,
		"float":    floatValue,
		"transfer": transferValue,
	}).Info("Calculating transfer value")

	// Call smart contract
	beneficiary := common.HexToAddress(transfer.Beneficiary)
	transaction, err := h.Contract.BuyTokensBTC(h.TxOpts, beneficiary, transferValue)
	if err != nil {
		return ctx.JsonError(err)
	}

	// Begin transaction
	tx := h.Database.Begin()

	// Update transfer
	transfer.Rate = h.Config.BtcEthFallbackRate
	transfer.TxIn = c.QueryParam("transaction_hash")
	transfer.TxOut = transaction.Hash().String()
	transfer.ValueIn = strconv.FormatInt(value, 10)
	transfer.ValueOut = transferValue.String()
	if err := h.Database.Save(&transfer).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Delete transfer
	if err := h.Database.Delete(transfer).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Delete address
	if err := h.Database.Delete(transfer.Address).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Decrease gap
	transfer.Address.Account.Gap = transfer.Address.Account.Gap - 1
	if err := h.Database.Save(transfer.Address.Account).Error; err != nil {
		tx.Rollback()
		return ctx.JsonError(err)
	}

	// Commit
	tx.Commit()

	return ctx.String(http.StatusOK, core.BlockchainInfoCallbackOk)
}
