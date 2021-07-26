package api

import (
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nsip/otf-config/config"
	"github.com/pkg/errors"
)

func Factory4NewUpdate(NewUpdate string) func(c echo.Context) error {

	mtx := &sync.Mutex{}

	return func(c echo.Context) error {
		defer mtx.Unlock()
		mtx.Lock()

		mNew := map[string]config.IEle{
			"NatsStreaming": new(config.NatsStreaming),
			"Nias3":         new(config.Nias3),
			"Benthos":       new(config.Benthos),
			"Reader":        new(config.Reader),
			"Align":         new(config.Align),
			"TxtClassifier": new(config.TxtClassifier),
			"Level":         new(config.Level),
			"Weight":        new(config.Weight),
			"Hub":           new(config.Hub),
		}

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

		var (
			proj    = c.QueryParam("project")
			failmsg = "Add/Update Config Failed "
			status  = http.StatusOK
			info    = "OTF Config Added/Updated "
			newcfg  = mNew[proj]
		)

		if err := c.Bind(newcfg); err != nil {
			status = http.StatusBadRequest
			info = errors.Wrap(err, "Bind Error - "+failmsg).Error()
			goto R
		}

		if err := newcfg.Validate(); err != nil {
			status = http.StatusBadRequest
			info = errors.Wrap(err, "Validate Error - "+failmsg+newcfg.GetName()).Error()
			goto R
		}

		switch NewUpdate {
		case "New", "new", "POST", "post":
			mGrp[proj].Add(newcfg)
		case "Update", "update", "PUT", "put":
			mGrp[proj].Update(newcfg.GetName(), newcfg)
		default:
			panic("Invalid 'NewUpdate' value")
		}

		if err := cfg.SaveToml(); err != nil {
			status = http.StatusInternalServerError
			info = errors.Wrap(err, "SaveToml Error - "+failmsg+newcfg.GetName()).Error()
			goto R
		}

		info += " @ " + newcfg.GetName()

		// Dispense
		time.Sleep(20 * time.Millisecond) // give time for config.toml updating
		cfg.Dispense(proj)

	R:
		return c.JSON(status, info)
	}
}
