package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nsip/otf-config/config"
	"github.com/pkg/errors"
)

func NewNatsStreamingCfg(c echo.Context) error {

	var (
		failmsg = "Adding Config Failed: "
		status  = http.StatusOK
		info    = "OTF Config Added: "
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

	cfg.NatsStreamings = append(cfg.NatsStreamings, *newcfg)

	if err := cfg.SaveToml(); err != nil {
		status = http.StatusInternalServerError
		info = errors.Wrap(err, failmsg+newcfg.Name).Error()
		goto R
	}

	info += newcfg.Name

R:
	return c.JSON(status, info)
}

func NewNias3Cfg(c echo.Context) error {

	var (
		failmsg = "Adding Config Failed: "
		status  = http.StatusOK
		info    = "OTF Config Added: "
	)

	newcfg := new(config.Nias3)

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

	cfg.Nias3s = append(cfg.Nias3s, *newcfg)

	if err := cfg.SaveToml(); err != nil {
		status = http.StatusInternalServerError
		info = errors.Wrap(err, failmsg+newcfg.Name).Error()
		goto R
	}

	info += newcfg.Name

R:
	return c.JSON(status, info)
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
