package main

import (
	"github.com/ShoppersShop/coinflip/core"
	"github.com/ShoppersShop/coinflip/handlers"
	echorelic "github.com/jessie-codes/echo-relic"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	newrelic "github.com/newrelic/go-agent"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Init config & Coinflip handler
	cfg := core.NewConfig("cf")
	coinflip := handlers.NewCoinflip(cfg)

	// Init echo
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())

	// Configure NewRelic if necessary
	config := newrelic.NewConfig(cfg.AppName, cfg.NewRelicLicenseKey)
	app, err := newrelic.NewApplication(config)
	if err != nil {
		log.Fatalf(core.ErrNewRelicAgent, err.Error())
	}
	e.Use(echorelic.Middleware(app))

	// Wire up custom context
	e.Use(func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := &core.CoinflipContext{c}
			return h(ctx)
		}
	})

	// Healthcheck route
	e.GET("/", coinflip.Healthcheck)

	// Stats feature
	if coinflip.HasFeature(core.FeatureStats) {
		g1 := e.Group("stats")
		g1.GET("", coinflip.StatsGet)
		g1.GET("/balance/:address", coinflip.BalanceGet)
	}

	// Whitelist feature
	if coinflip.HasFeature(core.FeatureWhitelist) {
		g2 := e.Group("whitelist")
		g2.GET("/:address", coinflip.WhitelistGet)
		g2.POST("", coinflip.WhitelistPost)
		g2.DELETE("", coinflip.WhitelistDelete)
	}

	// Blockchain.info feature
	if coinflip.HasFeature(core.FeatureBlockchain) {
		g3 := e.Group("blockchain")
		g3.POST("/receive", coinflip.BlockchainReceive)
		g3.GET("/callback/:invoice_id", coinflip.BlockchainCallback)
		g3.GET("/invoice/:id", coinflip.BlockchainInvoiceGet)
		g3.GET("/callback/logs", coinflip.BlockchainCallbackLogs)
		g3.GET("/gap/:xpub", coinflip.BlockchainGapCheck)
	}

	// Start server
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
