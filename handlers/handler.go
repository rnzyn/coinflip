package handlers

import (
	"github.com/ShoppersShop/coinflip/contracts"
	"github.com/ShoppersShop/coinflip/core"
	httpclient "github.com/ddliu/go-httpclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type Coinflip struct {
	Config   *core.Config
	TxOpts   *bind.TransactOpts
	Contract *contracts.TokenSale
}

func NewCoinflip(cfg *core.Config) *Coinflip {
	// Configure http client
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_DEBUG: cfg.HttpClientDebug,
		"Content-Type":       "application/json",
		"Accept":             "application/json",
	})

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(cfg.GethIPC)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %s", err)
	}

	// Instantiate the contract and display its name
	contract, err := contracts.NewTokenSale(common.HexToAddress(cfg.ContractAddress), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %s", err)
	}

	// Load private key
	ecdsaKey, err := crypto.HexToECDSA(cfg.GethPrivateKey)
	if err != nil {
		log.Fatalf("Failure: %v", err)
	}

	// Instantiate handler
	h := new(Coinflip)
	h.Config = cfg
	h.Contract = contract
	h.TxOpts = bind.NewKeyedTransactor(ecdsaKey)
	return h
}

func (h *Coinflip) HasFeature(feature string) bool {
	for _, item := range h.Config.Features {
		if feature == item {
			return true
		}
	}

	return false
}
