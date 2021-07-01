package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nsip/otf-config/config"
	"github.com/pkg/errors"
)

func UpdateNatsStreamingCfg(c echo.Context) error {

	var (
		failmsg = "Update Config Failed: "
		status  = http.StatusOK
		info    = "OTF Config Updated: "
	)

	newcfg := new(config.NatsStreaming)

	if err := c.Bind(newcfg); err != nil {
		status = http.StatusBadRequest
		info = errors.Wrap(err, failmsg).Error()
		goto R
	}

	if err := newcfg.Validate(); err != nil {
		status = http.StatusBadRequest
		info = errors.Wrap(err, failmsg+newcfg.Name).Error()
		goto R
	}

	for i, c := range cfg.NatsStreamings {
		if c.Name == newcfg.Name {
			cfg.NatsStreamings[i] = *newcfg
			break
		}
	}

	if err := cfg.SaveToml(); err != nil {
		status = http.StatusInternalServerError
		info = errors.Wrap(err, failmsg+newcfg.Name).Error()
		goto R
	}

	info += newcfg.Name

R:
	return c.JSON(status, info)
}

func UpdateNias3Cfg(c echo.Context) error {
	return nil
}

func UpdateBenthosCfg(c echo.Context) error {
	return nil
}

func UpdateReaderCfg(c echo.Context) error {
	return nil
}

func UpdateAlignCfg(c echo.Context) error {
	return nil
}

func UpdateTextclassifierCfg(c echo.Context) error {
	return nil
}

func UpdateLevelCfg(c echo.Context) error {
	return nil
}

func UpdateWeightCfg(c echo.Context) error {
	return nil
}

func UpdateHubCfg(c echo.Context) error {
	return nil
}
