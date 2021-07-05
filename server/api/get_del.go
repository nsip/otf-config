package api

import (
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/nsip/otf-config/config"
	"github.com/pkg/errors"
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
		m[k] = v.AllNames()
	}

	return c.JSON(http.StatusOK, m)
}

func Factory4GetDel(GetDel, proj string) func(c echo.Context) error {

	mtx := &sync.Mutex{}

	return func(c echo.Context) error {
		defer mtx.Unlock()
		mtx.Lock()

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

		name := c.QueryParam(cfgNameQuery)

		switch GetDel {
		case "Get", "GET", "get":
			return c.JSON(http.StatusOK, mGrp[proj].Get(name))
		case "Delete", "DELETE", "delete", "Del", "del":
			// log4del("In Delete")
			mGrp[proj].Delete(name)
			if err := cfg.SaveToml(); err != nil {
				info := errors.Wrap(err, "SaveToml Error - Delete Failed"+name).Error()
				return c.JSON(http.StatusInternalServerError, info)
			}
			return c.JSON(http.StatusOK, name+" deleted")
		default:
			panic("Invalid 'GetDel' value")
		}
	}
}
