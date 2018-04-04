package main

import (
	"github.com/ShoppersShop/coinflip/core"
	"github.com/ShoppersShop/coinflip/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Init config & coinflip
	cfg := core.NewConfig("coinflip")
	coinflip := core.NewCoinflip(cfg)

	// Init echo
	e := echo.New()
	e.HideBanner = true
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &core.CoinflipContext{c, coinflip}
			return h(ctx)
		}
	})

	// Configure middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Healthcheck route
	e.GET("/", handlers.Healthcheck)

	// Stats feature
	if coinflip.HasFeature("stats") {
		e.GET("/info", handlers.GetStats)
	}

	// Whitelist feature
	if coinflip.HasFeature("whitelist") {
		e.GET("/whitelist/:address", handlers.WhitelistGet)
		e.POST("/whitelist", handlers.WhitelistAdd)
		e.DELETE("/whitelist", handlers.WhitelistRemove)
	}

	// Bitcoin feature
	if coinflip.HasFeature("bitcoin") {
		e.POST("/bitcoin", handlers.BitcoinDonation)
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
