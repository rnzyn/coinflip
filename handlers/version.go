package handlers

import (
	"net/http"

	"github.com/ShoppersShop/coinflip/core"
	"github.com/labstack/echo"
)

func (h *Coinflip) Version(c echo.Context) error {
	ctx := c.(*core.CoinflipContext)
	return ctx.JSON(http.StatusOK, map[string]string{
		"git_commit":  h.Config.Build.GitCommit,
		"git_branch":  h.Config.Build.GitBranch,
		"git_state":   h.Config.Build.GitState,
		"git_summary": h.Config.Build.GitSummary,
		"build_date":  h.Config.Build.BuildDate,
	})
}
