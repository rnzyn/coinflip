package core

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Domain               string
	Port                 string
	Features             []string
	HttpClientDebug      bool
	GethIPC              string
	GethPrivateKey       string
	ContractAddress      string
	BlockchainInfoApiKey string
	BitcoinAccountXpub   string
	BtcEthFallbackRate   float64
}

func NewConfig(prefix string) *Config {
	// Load .env if exists
	godotenv.Load()

	// Viper instance
	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)
	viper.SetDefault("port", "3000")
	viper.SetDefault("http_client_debug", false)
	viper.SetDefault("features", "stats whitelist bitcoin")

	// Load configuration variables
	cfg := new(Config)
	cfg.Port = viper.GetString("port")
	cfg.Features = viper.GetStringSlice("features")
	cfg.HttpClientDebug = viper.GetBool("http_client_debug")

	// Fail fast
	cfg.Domain = viper.GetString("domain")
	if cfg.Domain == "" {
		log.Fatalf("Please set mandatory `CF_DOMAIN` configuration variable")
	}

	cfg.GethIPC = viper.GetString("geth_ipc")
	if cfg.GethIPC == "" {
		log.Fatalf("Please set mandatory `CF_GETH_IPC` configuration variable")
	}

	cfg.GethPrivateKey = viper.GetString("geth_private_key")
	if cfg.GethPrivateKey == "" {
		log.Fatalf("Please set mandatory `CF_GETH_PRIVATE_KEY` configuration variable")
	}

	cfg.ContractAddress = viper.GetString("contract_address")
	if cfg.ContractAddress == "" {
		log.Fatalf("Please set mandatory `CF_CONTRACT_ADDRESS` configuration variable")
	}

	cfg.BlockchainInfoApiKey = viper.GetString("blockchain_info_api_key")
	if cfg.BlockchainInfoApiKey == "" {
		log.Fatalf("Please set mandatory `CF_BLOCKCHAIN_INFO_API_KEY` configuration variable")
	}

	cfg.BitcoinAccountXpub = viper.GetString("bitcoin_account_xpub")
	if cfg.BitcoinAccountXpub == "" {
		log.Fatalf("Please set mandatory `CF_BITCOIN_ACCOUNT_XPUB` configuration variable")
	}

	cfg.BtcEthFallbackRate = viper.GetFloat64("btceth_fallback_rate")
	if cfg.BtcEthFallbackRate == 0 {
		log.Fatalf("Please set mandatory `CF_BTCETH_FALLBACK_RATE` configuration variable")
	}

	// Print configuration variables
	log.WithFields(log.Fields{
		"domain":               cfg.Domain,
		"port":                 cfg.Port,
		"http_client_debug":    cfg.HttpClientDebug,
		"features":             cfg.Features,
		"geth_ipc":             cfg.GethIPC,
		"contract_address":     cfg.ContractAddress,
		"bitcoin_xpub_key":     cfg.BitcoinAccountXpub,
		"btceth_fallback_rate": cfg.BtcEthFallbackRate,
	}).Info("Coinflip configuration")

	return cfg
}
