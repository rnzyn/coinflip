package core

import (
	"github.com/ShoppersShop/coinflip/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

type Coinflip struct {
	echo.Context
	auth     *bind.TransactOpts
	sale     *contracts.TokenSale
	features []string
}

func NewCoinflip(cfg *Config) *Coinflip {
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
	flip.auth = bind.NewKeyedTransactor(ecdsaKey)
	flip.sale = sale
	flip.features = cfg.Features
	return flip
}

func (flip *Coinflip) Auth() *bind.TransactOpts {
	return flip.auth
}

func (flip *Coinflip) Sale() *contracts.TokenSale {
	return flip.sale
}

func (flip *Coinflip) HasFeature(feature string) bool {
	for _, item := range flip.features {
		if feature == item {
			return true
		}
	}

	return false
}
