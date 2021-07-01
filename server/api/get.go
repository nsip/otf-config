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
	m := make(map[string][]string)

	for _, c := range cfg.NatsStreamings {
		m["NatsStreamings"] = append(m["NatsStreamings"], c.Name)
	}

	for _, c := range cfg.Nias3s {
		m["Nias3s"] = append(m["Nias3s"], c.Name)
	}

	for _, c := range cfg.Benthoses {
		m["Benthoses"] = append(m["Benthoses"], c.Name)
	}

	for _, c := range cfg.Readers {
		m["Readers"] = append(m["Readers"], c.Name)
	}

	for _, c := range cfg.Aligns {
		m["Aligns"] = append(m["Aligns"], c.Name)
	}

	for _, c := range cfg.TxtClassifiers {
		m["TxtClassifiers"] = append(m["TxtClassifiers"], c.Name)
	}

	for _, c := range cfg.Levels {
		m["Levels"] = append(m["Levels"], c.Name)
	}

	for _, c := range cfg.Weights {
		m["Weights"] = append(m["Weights"], c.Name)
	}

	for _, c := range cfg.Hubs {
		m["Hubs"] = append(m["Hubs"], c.Name)
	}

	return c.JSON(http.StatusOK, m)
}

func Factory4GetCfg(proj string) func(c echo.Context) error {

	m := map[string]interface{}{
		"NatsStreaming": cfg.NatsStreamings,
		"Nias3":         cfg.Nias3s,
		"Benthos":       cfg.Benthoses,
		"Reader":        cfg.Readers,
		"Align":         cfg.Aligns,
		"TxtClassifier": cfg.TxtClassifiers,
		"Level":         cfg.Levels,
		"Weight":        cfg.Weights,
		"Hub":           cfg.Hubs,
	}

	return func(c echo.Context) error {

		cname := c.QueryParam(cfgNameQuery)

		switch arr := m[proj].(type) {
		case []config.NatsStreaming:
			for _, e := range arr {
				if cname == e.GetName() {
					return c.JSON(http.StatusOK, e)
				}
			}
		case []config.Nias3:
			for _, e := range arr {
				if cname == e.GetName() {
					return c.JSON(http.StatusOK, e)
				}
			}
		case []config.Benthos:
			for _, e := range arr {
				if cname == e.GetName() {
					return c.JSON(http.StatusOK, e)
				}
			}
		case []config.Reader:
			for _, e := range arr {
				if cname == e.GetName() {
					return c.JSON(http.StatusOK, e)
				}
			}
		case []config.Align:
			for _, e := range arr {
				if cname == e.GetName() {
					return c.JSON(http.StatusOK, e)
				}
			}
		case []config.TxtClassifier:
			for _, e := range arr {
				if cname == e.GetName() {
					return c.JSON(http.StatusOK, e)
				}
			}
		case []config.Level:
			for _, e := range arr {
				if cname == e.GetName() {
					return c.JSON(http.StatusOK, e)
				}
			}
		case []config.Weight:
			for _, e := range arr {
				if cname == e.GetName() {
					return c.JSON(http.StatusOK, e)
				}
			}
		case []config.Hub:
			for _, e := range arr {
				if cname == e.GetName() {
					return c.JSON(http.StatusOK, e)
				}
			}
		}

		return nil
	}
}
