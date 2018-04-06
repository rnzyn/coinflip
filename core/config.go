package core

import (
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Domain               string
	Port                 string
	Features             []string
	HttpClientDebug      bool
	GethIpcPath          string
	EthPrivateKey        string
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
	viper.SetDefault(ConfigOptionPort, "3000")
	viper.SetDefault(ConfigOptionHttpClientDebug, false)
	viper.SetDefault(ConfigOptionFeatures, "stats whitelist bitcoin")

	// Load configuration variables
	cfg := new(Config)
	cfg.Port = viper.GetString(ConfigOptionPort)
	cfg.Features = viper.GetStringSlice(ConfigOptionFeatures)
	cfg.HttpClientDebug = viper.GetBool(ConfigOptionHttpClientDebug)

	// Fail fast
	cfg.Domain = viper.GetString(ConfigOptionDomain)
	if cfg.Domain == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionDomain)
		log.Fatalf("Please set mandatory `%s` configuration variable", option)
	}

	cfg.GethIpcPath = viper.GetString(ConfigOptionGethIpcPath)
	if cfg.GethIpcPath == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionGethIpcPath)
		log.Fatalf("Please set mandatory `%s` configuration variable", option)
	}

	cfg.EthPrivateKey = viper.GetString(ConfigOptionEthPrivateKey)
	if cfg.EthPrivateKey == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionEthPrivateKey)
		log.Fatalf("Please set mandatory `%s` configuration variable", option)
	}

	cfg.ContractAddress = viper.GetString(ConfigOptionContractAddress)
	if cfg.ContractAddress == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionContractAddress)
		log.Fatalf("Please set mandatory `%s` configuration variable", option)
	}

	cfg.BlockchainInfoApiKey = viper.GetString(ConfigOptionBlockchainInfoApiKey)
	if cfg.BlockchainInfoApiKey == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionBlockchainInfoApiKey)
		log.Fatalf("Please set mandatory `%s` configuration variable", option)
	}

	cfg.BitcoinAccountXpub = viper.GetString(ConfigOptionBitcoinAccountXpub)
	if cfg.BitcoinAccountXpub == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionBitcoinAccountXpub)
		log.Fatalf("Please set mandatory `%s` configuration variable", option)
	}

	cfg.BtcEthFallbackRate = viper.GetFloat64(ConfigOptionBtcEthFallbackRate)
	if cfg.BtcEthFallbackRate == 0 {
		option := strings.ToUpper(prefix + "_" + ConfigOptionDomain)
		log.Fatalf("Please set mandatory `%s` configuration variable", option)
	}

	// Print configuration variables
	log.WithFields(log.Fields{
		ConfigOptionDomain:             cfg.Domain,
		ConfigOptionPort:               cfg.Port,
		ConfigOptionHttpClientDebug:    cfg.HttpClientDebug,
		ConfigOptionFeatures:           cfg.Features,
		ConfigOptionGethIpcPath:        cfg.GethIpcPath,
		ConfigOptionContractAddress:    cfg.ContractAddress,
		ConfigOptionBitcoinAccountXpub: cfg.BitcoinAccountXpub,
		ConfigOptionBtcEthFallbackRate: cfg.BtcEthFallbackRate,
	}).Info("Coinflip configuration")

	return cfg
}
