package core

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	AppName              string
	Debug                bool
	Database             string
	Domain               string
	Port                 string
	Features             []string
	GethIpcPath          string
	EthPrivateKey        string
	EthSaleContract      string
	EthTokenContract     string
	BlockchainInfoApiKey string
	BtcEthFallbackRate   decimal.Decimal
	NewRelicLicenseKey   string
}

func NewConfig(prefix string) *Config {
	// Load .env if exists
	godotenv.Load()

	// Viper instance
	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)
	viper.SetDefault(ConfigDefaultAppName, ConfigDefaultAppName)
	viper.SetDefault(ConfigOptionPort, ConfigDefaultPort)
	viper.SetDefault(ConfigOptionDebug, ConfigDefaultDebug)
	viper.SetDefault(ConfigOptionFeatures, ConfigDefaultFeatures)

	// Load configuration variables
	cfg := new(Config)
	cfg.AppName = viper.GetString(ConfigOptionAppName)
	cfg.Port = viper.GetString(ConfigOptionPort)
	cfg.Features = viper.GetStringSlice(ConfigOptionFeatures)
	cfg.Debug = viper.GetBool(ConfigOptionDebug)
	cfg.NewRelicLicenseKey = viper.GetString(ConfigOptionNewRelicLicenseKey)

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

	cfg.EthSaleContract = viper.GetString(ConfigOptionEthSaleContract)
	if cfg.EthSaleContract == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionEthSaleContract)
		log.Fatalf(ErrConfigOptionRequired, option)
	}

	cfg.EthTokenContract = viper.GetString(ConfigOptionEthTokenContract)
	if cfg.EthTokenContract == "" {
		option := strings.ToUpper(prefix + "_" + ConfigOptionEthTokenContract)
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
		ConfigOptionAppName:            cfg.AppName,
		ConfigOptionDomain:             cfg.Domain,
		ConfigOptionPort:               cfg.Port,
		ConfigOptionDebug:              cfg.Debug,
		ConfigOptionFeatures:           cfg.Features,
		ConfigOptionGethIpcPath:        cfg.GethIpcPath,
		ConfigOptionEthSaleContract:    cfg.EthSaleContract,
		ConfigOptionEthTokenContract:   cfg.EthTokenContract,
		ConfigOptionBtcEthFallbackRate: cfg.BtcEthFallbackRate,
	}).Info("Coinflip configuration")

	cfg.BtcEthFallbackRate = fallbackRate
	return cfg
}
