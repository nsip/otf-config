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

func NatsStreamingCfg(c echo.Context) error {
	cname := c.QueryParam(cfgNameQuery)
	for _, cr := range cfg.NatsStreamings {
		if cr.Name == cname {
			return c.JSON(http.StatusOK, cr)
		}
	}
	return nil
}

func Nias3Cfg(c echo.Context) error {
	cname := c.QueryParam(cfgNameQuery)
	for _, cr := range cfg.Nias3s {
		if cr.Name == cname {
			return c.JSON(http.StatusOK, cr)
		}
	}
	return nil
}

func BenthosCfg(c echo.Context) error {
	cname := c.QueryParam(cfgNameQuery)
	for _, cr := range cfg.Benthoses {
		if cr.Name == cname {
			return c.JSON(http.StatusOK, cr)
		}
	}
	return nil
}

///

func ReaderCfg(c echo.Context) error {
	cname := c.QueryParam(cfgNameQuery)
	for _, cr := range cfg.Readers {
		if cr.Name == cname {
			return c.JSON(http.StatusOK, cr)
		}
	}
	return nil
}

func AlignCfg(c echo.Context) error {
	cname := c.QueryParam(cfgNameQuery)
	for _, cr := range cfg.Aligns {
		if cr.Name == cname {
			return c.JSON(http.StatusOK, cr)
		}
	}
	return nil
}

func TextclassifierCfg(c echo.Context) error {
	cname := c.QueryParam(cfgNameQuery)
	for _, cr := range cfg.TxtClassifiers {
		if cr.Name == cname {
			return c.JSON(http.StatusOK, cr)
		}
	}
	return nil
}

func LevelCfg(c echo.Context) error {
	cname := c.QueryParam(cfgNameQuery)
	for _, cr := range cfg.Levels {
		if cr.Name == cname {
			return c.JSON(http.StatusOK, cr)
		}
	}
	return nil
}

func WeightCfg(c echo.Context) error {
	cname := c.QueryParam(cfgNameQuery)
	for _, cr := range cfg.Weights {
		if cr.Name == cname {
			return c.JSON(http.StatusOK, cr)
		}
	}
	return nil
}

func HubCfg(c echo.Context) error {
	cname := c.QueryParam(cfgNameQuery)
	for _, cr := range cfg.Hubs {
		if cr.Name == cname {
			return c.JSON(http.StatusOK, cr)
		}
	}
	return nil
}
