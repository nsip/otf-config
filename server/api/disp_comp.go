package api

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	md "github.com/digisan/data-drawing/markdown"
	"github.com/labstack/echo/v4"
	"github.com/nsip/otf-config/config"
	"github.com/pkg/errors"
)

func Dispense(c echo.Context) error {
	log4get("dispensing...")
	cfg = config.GetConfig("../config.toml", "./config.toml")
	cfg.Dispense()
	return c.JSON(http.StatusOK, "Dispensed")
}

func Composite(c echo.Context) error {

	log4post("compositing...")

	var (
		failmsg = "Compositing Failed "
		status  = http.StatusOK
		info    = "Composited"
	)

	type TableCols struct {
		ExePathGrp *[]interface{}
		ArgsGrp    *[]interface{}
		DelayGrp   *[]interface{}
		Enabled    *[]interface{}
	}

	cols := &TableCols{}
	if err := c.Bind(cols); err != nil {
		status = http.StatusBadRequest
		info = errors.Wrap(err, "Bind Error - "+failmsg).Error()
		goto R
	}

	// spew.Dump(cols)

	md.MDTable(cols, "./test.md")

R:
	return c.JSON(status, info)
}
