package core

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Debug                bool
	Database             string
	Domain               string
	Port                 string
	Features             []string
	GethIpcPath          string
	EthPrivateKey        string
	ContractAddress      string
	BlockchainInfoApiKey string
	BtcEthFallbackRate   decimal.Decimal
}

func NewConfig(prefix string) *Config {
	// Load .env if exists
	godotenv.Load()

	// Viper instance
	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)
	viper.SetDefault(ConfigOptionPort, ConfigDefaultPort)
	viper.SetDefault(ConfigOptionDebug, ConfigDefaultDebug)
	viper.SetDefault(ConfigOptionFeatures, ConfigDefaultFeatures)

	// Load configuration variables
	cfg := new(Config)
	cfg.Port = viper.GetString(ConfigOptionPort)
	cfg.Features = viper.GetStringSlice(ConfigOptionFeatures)
	cfg.Debug = viper.GetBool(ConfigOptionDebug)

	// Fail fast
	cfg.Database = viper.GetString(ConfigOptionDatabase)
	if cfg.Database == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionDatabase)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.Domain = viper.GetString(ConfigOptionDomain)
	if cfg.Domain == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionDomain)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.GethIpcPath = viper.GetString(ConfigOptionGethIpcPath)
	if cfg.GethIpcPath == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionGethIpcPath)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.EthPrivateKey = viper.GetString(ConfigOptionEthPrivateKey)
	if cfg.EthPrivateKey == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionEthPrivateKey)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.ContractAddress = viper.GetString(ConfigOptionContractAddress)
	if cfg.ContractAddress == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionContractAddress)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.BlockchainInfoApiKey = viper.GetString(ConfigOptionBlockchainInfoApiKey)
	if cfg.BlockchainInfoApiKey == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionBlockchainInfoApiKey)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	fallbackRate, err := decimal.NewFromString(viper.GetString(ConfigOptionBtcEthFallbackRate))
	if err != nil {
		log.Fatal(err)
	}

	// Print configuration variables
	log.WithFields(log.Fields{
		ConfigOptionDomain:             cfg.Domain,
		ConfigOptionPort:               cfg.Port,
		ConfigOptionDebug:              cfg.Debug,
		ConfigOptionDatabase:           cfg.Database,
		ConfigOptionFeatures:           cfg.Features,
		ConfigOptionGethIpcPath:        cfg.GethIpcPath,
		ConfigOptionContractAddress:    cfg.ContractAddress,
		ConfigOptionBtcEthFallbackRate: cfg.BtcEthFallbackRate,
	}).Info("Coinflip configuration")

	cfg.BtcEthFallbackRate = fallbackRate
	return cfg
}
