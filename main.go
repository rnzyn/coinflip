package main

import (
	"github.com/ShoppersShop/coinflip/core"
	"github.com/ShoppersShop/coinflip/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Init config & Coinflip handler
	cfg := core.NewConfig("cf")
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
		g1 := e.Group("stats")
		g1.GET("", coinflip.StatsGet)
	}

	// Whitelist feature
	if coinflip.HasFeature("whitelist") {
		g2 := e.Group("whitelist")
		g2.GET("/:address", coinflip.WhitelistGet)
		g2.POST("", coinflip.WhitelistPost)
		g2.DELETE("", coinflip.WhitelistDelete)
	}

	// Blockchain.info feature
	if coinflip.HasFeature("blockchain") {
		g3 := e.Group("blockchain")
		g3.POST("/receive", coinflip.BlockchainReceive)
		g3.GET("/callback", coinflip.BlockchainCallback)
		g3.GET("/invoice/:id", coinflip.BlockchainInvoiceGet)
		g3.GET("/callback/logs", coinflip.BlockchainCallbackLogs)
		g3.GET("/gap/:xpub", coinflip.BlockchainGapCheck)
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
