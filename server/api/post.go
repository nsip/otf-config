package api

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
)

func NewNatsStreamingCfg(c echo.Context) error {

	ns := new(struct {
		CfgName string `json:"name"`
		Path    string `json:"path"`
	})
	if err := c.Bind(ns); err != nil {
		log4post("%v", err)
	}
	spew.Dump(ns)

	// CANNOT Fetch by c.FormValue
	// cfg.NatsStreamings = append(cfg.NatsStreamings, struct {
	// 	CfgName string `json:"name"`
	// 	Path    string `json:"path"`
	// }{
	// 	CfgName: c.FormValue("name"),
	// 	Path:    c.FormValue("path"),
	// })
	// log4post("%s - %s", c.FormValue("name"), c.FormValue("path"))

	return c.JSON(http.StatusOK, "POST to NewNS")
}

func NewNias3Cfg(c echo.Context) error {

	ns := new(struct {
		CfgName string `json:"name"`
		Path    string `json:"path"`
	})
	if err := c.Bind(ns); err != nil {
		log4post("%v", err)
	}
	spew.Dump(ns)

	return c.JSON(http.StatusOK, "POST to NewNias3")
}

func NewBenthosCfg(c echo.Context) error {
	return nil
}

func NewReaderCfg(c echo.Context) error {
	return nil
}

func NewAlignCfg(c echo.Context) error {
	return nil
}

func NewTextclassifierCfg(c echo.Context) error {
	return nil
}

func NewLevelCfg(c echo.Context) error {
	return nil
}

func NewWeightCfg(c echo.Context) error {
	return nil
}

func NewHubCfg(c echo.Context) error {
	return nil
}
