package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nsip/otf-config/config"
	"github.com/pkg/errors"
)

func Factory4NewCfg(proj string) func(c echo.Context) error {

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

		var (
			failmsg = "Add Config Failed: "
			status  = http.StatusOK
			info    = "OTF Config Added: "
			newcfg  = mNew[proj]
		)

		if err := c.Bind(newcfg); err != nil {
			status = http.StatusBadRequest
			info = errors.Wrap(err, failmsg).Error()
			goto R
		}

		if err := newcfg.Validate(); err != nil {
			status = http.StatusBadRequest
			info = errors.Wrap(err, failmsg+newcfg.GetName()).Error()
			goto R
		}

		mGrp[proj].AddElem(newcfg)

		if err := cfg.SaveToml(); err != nil {
			status = http.StatusInternalServerError
			info = errors.Wrap(err, failmsg+newcfg.GetName()).Error()
			goto R
		}

		info += newcfg.GetName()

	R:
		return c.JSON(status, info)
	}
}
