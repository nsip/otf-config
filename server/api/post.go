package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nsip/otf-config/config"
	"github.com/pkg/errors"
)

func Factory4NewCfg(proj string) func(c echo.Context) error {

	return func(c echo.Context) error {

		var newcfg config.IValidate

		switch proj {
		case "NatsStreaming":
			newcfg = new(config.NatsStreaming)
		case "Nias3":
			newcfg = new(config.Nias3)
		case "Benthos":
			newcfg = new(config.Benthos)
		case "Reader":
			newcfg = new(config.Reader)
		case "Align":
			newcfg = new(config.Align)
		case "TxtClassifier":
			newcfg = new(config.TxtClassifier)
		case "Level":
			newcfg = new(config.Level)
		case "Weight":
			newcfg = new(config.Weight)
		case "Hub":
			newcfg = new(config.Hub)
		}

		var (
			failmsg = "Add Config Failed: "
			status  = http.StatusOK
			info    = "OTF Config Added: "
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

		switch proj {
		case "NatsStreaming":
			cfg.NatsStreamings = append(cfg.NatsStreamings, *(newcfg).(*config.NatsStreaming))
		case "Nias3":
			cfg.Nias3s = append(cfg.Nias3s, *(newcfg).(*config.Nias3))
		case "Benthos":
			cfg.Benthoses = append(cfg.Benthoses, *(newcfg).(*config.Benthos))
		case "Reader":
			cfg.Readers = append(cfg.Readers, *(newcfg).(*config.Reader))
		case "Align":
			cfg.Aligns = append(cfg.Aligns, *(newcfg).(*config.Align))
		case "TxtClassifier":
			cfg.TxtClassifiers = append(cfg.TxtClassifiers, *(newcfg).(*config.TxtClassifier))
		case "Level":
			cfg.Levels = append(cfg.Levels, *(newcfg).(*config.Level))
		case "Weight":
			cfg.Weights = append(cfg.Weights, *(newcfg).(*config.Weight))
		case "Hub":
			cfg.Hubs = append(cfg.Hubs, *(newcfg).(*config.Hub))
		}

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
