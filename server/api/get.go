package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nsip/otf-config/config"
)

func AllCfgGrp(c echo.Context) error {

	cfg = config.GetConfig("../config.toml", "./config.toml")
	return c.JSON(http.StatusOK, WantedFieldsByType(cfg, "[]struct"))
}

func AllCfgItems(c echo.Context) error {

	cfg = config.GetConfig("../config.toml", "./config.toml")

	mGrp := map[string]config.IGrp{
		"NatsStreaming": &cfg.NatsStreamings,
		"Nias3":         &cfg.Nias3s,
		"Benthos":       &cfg.Benthoses,
		"Reader":        &cfg.Readers,
		"Align":         &cfg.Aligns,
		"TxtClassifier": &cfg.TxtClassifiers,
		"Level":         &cfg.Levels,
		"Weight":        &cfg.Weights,
		"Hub":           &cfg.Hubs,
	}

	m := make(map[string][]string)
	for k, v := range mGrp {
		m[k] = v.GetAllNames()
	}

	return c.JSON(http.StatusOK, m)
}

func Factory4GetCfg(proj string) func(c echo.Context) error {

	return func(c echo.Context) error {

		mGrp := map[string]config.IGrp{
			"NatsStreaming": &cfg.NatsStreamings,
			"Nias3":         &cfg.Nias3s,
			"Benthos":       &cfg.Benthoses,
			"Reader":        &cfg.Readers,
			"Align":         &cfg.Aligns,
			"TxtClassifier": &cfg.TxtClassifiers,
			"Level":         &cfg.Levels,
			"Weight":        &cfg.Weights,
			"Hub":           &cfg.Hubs,
		}

		cname := c.QueryParam(cfgNameQuery)

		return c.JSON(http.StatusOK, mGrp[proj].GetElem(cname))
	}
}
