package core

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Port     string
	Contract string
	IPC      string
	Features []string
	Key      string
}

func NewConfig(prefix string) *Config {
	// Load .env if exists
	godotenv.Load()

	// Viper instance
	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)
	viper.SetDefault("port", "3000")
	viper.SetDefault("features", "stats whitelist bitcoin")

	// Load configuration variables
	cfg := new(Config)
	cfg.Port = viper.GetString("port")
	cfg.Contract = viper.GetString("contract")
	cfg.IPC = viper.GetString("ipc")
	cfg.Features = viper.GetStringSlice("features")
	cfg.Key = viper.GetString("key")

	// Fail fast
	if cfg.IPC == "" || cfg.Contract == "" || cfg.Key == "" {
		log.Fatal("Please set all required variables before running")
	}

	// Print configuration variables
	log.WithFields(log.Fields{
		"port":     cfg.Port,
		"contract": cfg.Contract,
		"ipc":      cfg.IPC,
		"features": cfg.Features,
	}).Info("Coinflip configuration")

	return cfg
}
