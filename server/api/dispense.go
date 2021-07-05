package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nsip/otf-config/config"
)

func Dispense(c echo.Context) error {
	log4get("dispensing...")
	cfg = config.GetConfig("../config.toml", "./config.toml")
	cfg.Dispense()
	return c.JSON(http.StatusOK, "Dispensed")
}
