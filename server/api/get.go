package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nsip/otf-config/config"
)

var (
	cfg          = config.OtfCfg
	cfgNameQuery = "cfgName"
)

func ReaderCfg(c echo.Context) error {
	cname := c.QueryParam(cfgNameQuery)
	for _, cr := range cfg.Readers {
		if cr.CfgName == cname {
			return c.JSON(http.StatusOK, cr)
		}
	}
	return nil
}

func AlignCfg(c echo.Context) error {
	return nil
}

func TextclassifierCfg(c echo.Context) error {
	return nil
}

func LevelCfg(c echo.Context) error {
	return nil
}

func WeightCfg(c echo.Context) error {
	return nil
}

func HubCfg(c echo.Context) error {
	return nil
}
