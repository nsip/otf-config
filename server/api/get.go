package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AllCfgGrp(c echo.Context) error {
	return c.JSON(http.StatusOK, WantedFieldsByType(cfg, "[]struct"))
}

func AllCfgItems(c echo.Context) error {

	m := make(map[string][]string)

	for _, c := range cfg.NatsStreamings {
		m["NatsStreamings"] = append(m["NatsStreamings"], c.CfgName)
	}

	for _, c := range cfg.Nias3s {
		m["Nias3s"] = append(m["Nias3s"], c.CfgName)
	}

	for _, c := range cfg.Benthoses {
		m["Benthoses"] = append(m["Benthoses"], c.CfgName)
	}

	for _, c := range cfg.Readers {
		m["Readers"] = append(m["Readers"], c.CfgName)
	}

	for _, c := range cfg.Aligns {
		m["Aligns"] = append(m["Aligns"], c.CfgName)
	}

	for _, c := range cfg.TxtClassifiers {
		m["TxtClassifiers"] = append(m["TxtClassifiers"], c.CfgName)
	}

	for _, c := range cfg.Levels {
		m["Levels"] = append(m["Levels"], c.CfgName)
	}

	for _, c := range cfg.Weights {
		m["Weights"] = append(m["Weights"], c.CfgName)
	}

	for _, c := range cfg.Hubs {
		m["Hubs"] = append(m["Hubs"], c.CfgName)
	}

	return c.JSON(http.StatusOK, m)
}

func ReaderCfg(c echo.Context) error {
	cname := c.QueryParam(cfgNameQuery)
	for _, cr := range cfg.Readers {
		if cr.CfgName == cname {
			return c.JSON(http.StatusOK, cr)
		}
	}
	return nil
}

func AlignCfg(c echo.Context) error {
	return nil
}

func TextclassifierCfg(c echo.Context) error {
	return nil
}

func LevelCfg(c echo.Context) error {
	return nil
}

func WeightCfg(c echo.Context) error {
	return nil
}

func HubCfg(c echo.Context) error {
	return nil
}
