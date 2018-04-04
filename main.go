package main

import (
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	// Load dotenv
	godotenv.Load()

	// Viper instance
	viper.AutomaticEnv()
	viper.SetEnvPrefix("coinflip")
	viper.SetDefault("port", "3000")
	viper.SetDefault("ipc", "")
	viper.SetDefault("contract", "")
	viper.SetDefault("features", "stats whitelist bitcoin")
	viper.SetDefault("key", "")

	// Load configuration variables
	port := viper.GetString("port")
	contract := viper.GetString("contract")
	ipc := viper.GetString("ipc")
	features := viper.GetStringSlice("features")
	key := viper.GetString("key")

	// Print configuration variables
	log.WithFields(log.Fields{
		"port":     port,
		"contract": contract,
		"ipc":      ipc,
		"features": features,
	}).Info("Coinflip configuration")

	// Create an IPC based RPC connection to a remote node
	conn, err := ethclient.Dial(ipc)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %s", err)
	}

	// Instantiate the contract and display its name
	sale, err := NewTokenSale(common.HexToAddress(contract), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate contract: %s", err)
	}

	// Load private key
	ecdsaKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Fatalf("Failure: %v", err)
	}

	// Prepare transaction signer
	auth := bind.NewKeyedTransactor(ecdsaKey)

	// Echo instance
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Healthcheck
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	// Stats
	if isInSlice("stats", features) {
		e.GET("/stats", func(c echo.Context) error {
			active, err := sale.IsActiveSale(nil)
			availableUnits, err := sale.AvailableUnits(nil)
			basePrice, err := sale.BasePrice(nil)
			discountPrice, err := sale.DiscountPrice(nil)
			duration, err := sale.Duration(nil)
			minPayment, err := sale.MinPayment(nil)
			startTime, err := sale.StartTime(nil)
			unitsSold, err := sale.UnitsSold(nil)
			weiReceived, err := sale.WeiReceived(nil)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

			return c.JSON(http.StatusOK, map[string]interface{}{
				"active":         active,
				"availableUnits": availableUnits,
				"basePrice":      basePrice,
				"discountPrice":  discountPrice,
				"duration":       duration,
				"minPayment":     minPayment,
				"startTime":      startTime,
				"unitsSold":      unitsSold,
				"weiReceived":    weiReceived,
			})
		})
	}

	// Whitelist feature
	if isInSlice("whitelist", features) {
		e.GET("/whitelist/:address", func(c echo.Context) error {
			result, err := sale.WhitelistCheck(nil, common.HexToAddress(c.Param("address")))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

			return c.JSON(http.StatusOK, map[string]interface{}{
				"whitelisted": result,
			})
		})

		e.POST("/whitelist", func(c echo.Context) error {
			// Parse payload
			payload := new(WhitelistAdd)
			if err = c.Bind(payload); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

			// Convert raw input to Ethereum addresses
			addresses := []common.Address{}
			for _, address := range payload.Addresses {
				addresses = append(addresses, common.HexToAddress(address))
			}

			// Send transaction
			options := &bind.TransactOpts{From: auth.From, Signer: auth.Signer}
			transaction, err := sale.WhitelistAdd(options, addresses)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

			// Return response
			return c.JSON(http.StatusOK, map[string]interface{}{
				"txHash": transaction.Hash().String(),
			})
		})

		e.DELETE("/whitelist", func(c echo.Context) error {
			// Parse payload
			payload := new(WhitelistDelete)
			if err = c.Bind(payload); err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

			// Convert raw input to Ethereum addresses
			addresses := []common.Address{}
			for _, address := range payload.Addresses {
				addresses = append(addresses, common.HexToAddress(address))
			}

			// Send transaction
			options := &bind.TransactOpts{From: auth.From, Signer: auth.Signer}
			transaction, err := sale.WhitelistRemove(options, addresses)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			}

			// Return response
			return c.JSON(http.StatusOK, map[string]interface{}{
				"txHash": transaction.Hash().String(),
			})
		})
	}

	// Bitcoin feature
	if isInSlice("bitcoin", features) {
		e.POST("/bitcoin", func(c echo.Context) error {
			return c.JSON(http.StatusOK, map[string]interface{}{"error": "TBD"})
		})
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
