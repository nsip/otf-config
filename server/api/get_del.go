package api

import (
	"net/http"
	"sync"
	"time"

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

func Factory4GetDel(GetDel string) func(c echo.Context) error {

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

		proj := c.QueryParam(pNameProject)
		cname := c.QueryParam(pNameCfgName)

		switch GetDel {
		case "Get", "GET", "get":
			// log4del("In GET")

			return c.JSON(http.StatusOK, mGrp[proj].Get(cname))

		case "Delete", "DELETE", "delete", "Del", "del":
			// log4del("In DELETE")

			// Withdraw 1)
			mGrp[proj].Withdraw()
			time.Sleep(20 * time.Millisecond) // give 20 Millisecond for withdrawing

			mGrp[proj].Delete(cname)
			if err := cfg.SaveToml(); err != nil {
				info := errors.Wrap(err, "SaveToml Error - Delete Failed"+cname).Error()
				return c.JSON(http.StatusInternalServerError, info)
			}

			// Withdraw 2)
			time.Sleep(20 * time.Millisecond) // give 20 Millisecond for toml updating
			cfg.Dispense(proj)

			return c.JSON(http.StatusOK, cname+" deleted")
		default:
			panic("Invalid 'GetDel' value")
		}
	}
}
