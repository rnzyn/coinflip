package main

import (
	"github.com/ShoppersShop/coinflip/core"
	"github.com/ShoppersShop/coinflip/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Init config & Coinflip handler
	cfg := core.NewConfig("coinflip")
	coinflip := handlers.NewCoinflip(cfg)

	// Init echo
	e := echo.New()
	e.HideBanner = true
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &core.CoinflipContext{c}
			return h(ctx)
		}
	})

	// Configure middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Healthcheck route
	e.GET("/", coinflip.Healthcheck)

	// Stats feature
	if coinflip.HasFeature("stats") {
		e.GET("/stats", coinflip.StatsGet)
	}

	// Whitelist feature
	if coinflip.HasFeature("whitelist") {
		e.GET("/whitelist/:address", coinflip.WhitelistGet)
		e.POST("/whitelist", coinflip.WhitelistPost)
		e.DELETE("/whitelist", coinflip.WhitelistDelete)
	}

	// Bitcoin feature
	if coinflip.HasFeature("bitcoin") {
		e.POST("/bitcoin", coinflip.BitcoinDonation)
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
