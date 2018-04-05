package handlers

import (
	"github.com/ShoppersShop/coinflip/contracts"
	"github.com/ShoppersShop/coinflip/core"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type Coinflip struct {
	TxOpts   *bind.TransactOpts
	Contract *contracts.TokenSale
	Features []string
}

func NewCoinflip(cfg *core.Config) *Coinflip {
	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(cfg.IPC)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %s", err)
	}

	// Instantiate the contract and display its name
	sale, err := contracts.NewTokenSale(common.HexToAddress(cfg.Contract), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %s", err)
	}

	// Load private key
	ecdsaKey, err := crypto.HexToECDSA(cfg.Key)
	if err != nil {
		log.Fatalf("Failure: %v", err)
	}

	// Instantiate context
	flip := new(Coinflip)
	flip.Contract = sale
	flip.Features = cfg.Features
	flip.TxOpts = bind.NewKeyedTransactor(ecdsaKey)
	return flip
}

func (h *Coinflip) HasFeature(feature string) bool {
	for _, item := range h.Features {
		if feature == item {
			return true
		}
	}

	return false
}
