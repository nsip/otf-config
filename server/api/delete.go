package api

import (
	"github.com/labstack/echo/v4"
)

func Factory4Delete(proj string) func(c echo.Context) error {

	return func(c echo.Context) error {

		// mGrp := map[string]config.IGrp{
		// 	"NatsStreaming": &cfg.NatsStreamings,
		// 	"Nias3":         &cfg.Nias3s,
		// 	"Benthos":       &cfg.Benthoses,
		// 	"Reader":        &cfg.Readers,
		// 	"Align":         &cfg.Aligns,
		// 	"TxtClassifier": &cfg.TxtClassifiers,
		// 	"Level":         &cfg.Levels,
		// 	"Weight":        &cfg.Weights,
		// 	"Hub":           &cfg.Hubs,
		// }

		return nil
	}
}

// func DelNatsStreamingCfg(c echo.Context) error {
// 	return nil
// }

// func DelNias3Cfg(c echo.Context) error {
// 	return nil
// }

// func DelBenthosCfg(c echo.Context) error {
// 	return nil
// }

// func DelReaderCfg(c echo.Context) error {
// 	return nil
// }

// func DelAlignCfg(c echo.Context) error {
// 	return nil
// }

// func DelTextclassifierCfg(c echo.Context) error {
// 	return nil
// }

// func DelLevelCfg(c echo.Context) error {
// 	return nil
// }

// func DelWeightCfg(c echo.Context) error {
// 	return nil
// }

// func DelHubCfg(c echo.Context) error {
// 	return nil
// }
