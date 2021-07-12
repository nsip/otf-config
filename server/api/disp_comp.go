package api

import (
	"net/http"
	"path/filepath"
	"strings"

	md "github.com/digisan/data-drawing/markdown"
	"github.com/digisan/gotk/filedir"
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

	log4post("%s - %s - %s", "compositing...", c.QueryParam("name"), c.QueryParam("exepath"))

	var (
		failmsg = "Compositing Failed "
		status  = http.StatusOK
		info    = "Composited"
	)

	type TableCols struct {
		Indices    *[]interface{}
		Kinds      *[]interface{}
		ExePathGrp *[]interface{}
		ArgsGrp    *[]interface{}
		DelayGrp   *[]interface{}
		EnabledGrp *[]interface{}
	}

	var (
		cols   = &TableCols{}
		newhub = &config.Hub{
			Name:   "hub_config",            // single element, hard coded here
			Path:   c.QueryParam("exepath"), // hub exe path
			ComTbl: c.QueryParam("name"),    // composite table name
		}
		path, _ = filedir.AbsPath(newhub.Path, false)
		destTbl = filepath.Join(filepath.Dir(path), strings.TrimSuffix(newhub.ComTbl, ".md")+".md")
	)

	if err := c.Bind(cols); err != nil {
		status = http.StatusBadRequest
		info = errors.Wrap(err, "Bind Error - "+failmsg).Error()
		goto R
	}

	if err := newhub.Validate(); err != nil {
		status = http.StatusBadRequest
		info = errors.Wrap(err, "Validate Error - "+failmsg+newhub.GetName()).Error()
		goto R
	}

	cfg.Hubs.Update(newhub.Name, newhub)

	if err := cfg.SaveToml(); err != nil {
		status = http.StatusInternalServerError
		info = errors.Wrap(err, "SaveToml Error - "+failmsg+newhub.GetName()).Error()
		goto R
	}

	md.MDTable(cols, destTbl)

R:
	return c.JSON(status, info)
}
