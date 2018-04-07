package handlers

import (
	"github.com/ShoppersShop/coinflip/contracts"
	"github.com/ShoppersShop/coinflip/core"
	"github.com/ShoppersShop/coinflip/models"
	httpclient "github.com/ddliu/go-httpclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	log "github.com/sirupsen/logrus"
)

type Coinflip struct {
	Config   *core.Config
	Contract *contracts.TokenSale
	Database *gorm.DB
	TxOpts   *bind.TransactOpts
}

func NewCoinflip(cfg *core.Config) *Coinflip {
	// Configure http client
	httpclient.Defaults(httpclient.Map{
		httpclient.OPT_DEBUG: cfg.Debug,
		"Content-Type":       "application/json",
		"Accept":             "application/json",
	})

	// Configure database connection
	db, err := gorm.Open("postgres", cfg.Database)
	if err != nil {
		log.Fatalf(core.ErrDbConnFailure, err)
	}

	// Configure database handler
	db.LogMode(cfg.Debug)

	// Auto-migrate
	db.AutoMigrate(&models.Account{}, &models.Address{}, &models.Transfer{})

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(cfg.GethIpcPath)
	if err != nil {
		log.Fatalf(core.ErrGethConnFailrue, err)
	}

	// Instantiate the contract and display its name
	contract, err := contracts.NewTokenSale(common.HexToAddress(cfg.ContractAddress), conn)
	if err != nil {
		log.Fatalf(core.ErrContractInit, err)
	}

	// Load private key
	ecdsaKey, err := crypto.HexToECDSA(cfg.EthPrivateKey)
	if err != nil {
		log.Fatalf(core.ErrPrivateKeyInit, err)
	}

	// Instantiate handler
	h := new(Coinflip)
	h.Config = cfg
	h.Contract = contract
	h.Database = db
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
