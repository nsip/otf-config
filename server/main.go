package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/davecgh/go-spew/spew"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nsip/otf-config/config"
	"github.com/nsip/otf-config/server/api"
)

func hostPage() func() {

	cfg := config.GetConfig("../config.toml", "./config.toml")

	return func() {
		e := echo.New()
		e.Use(middleware.Gzip())
		e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		})) // allow cors requests during testing
		e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
			Index:  "index.html",
			Root:   cfg.PageFolder,
			Browse: true,
			HTML5:  true,
		}))

		e.GET("/allgrp", api.AllCfgGrp)
		e.GET("/allitems", api.AllCfgItems)

		grp := e.Group("/otf-config")
		// grp.Use(middleware.JWT("demoSecret"))

		// Search Config
		grp.GET("/natsstreaming", api.Factory4GetCfg("NatsStreaming"))
		grp.GET("/nias3", api.Factory4GetCfg("Nias3"))
		grp.GET("/benthos", api.Factory4GetCfg("Benthos"))
		grp.GET("/reader", api.Factory4GetCfg("Reader"))
		grp.GET("/align", api.Factory4GetCfg("Align"))
		grp.GET("/textclassifier", api.Factory4GetCfg("TxtClassifier"))
		grp.GET("/level", api.Factory4GetCfg("Level"))
		grp.GET("/weight", api.Factory4GetCfg("Weight"))
		grp.GET("/hub", api.Factory4GetCfg("Hub"))

		// New Config
		grp.POST("/natsstreaming", api.Factory4NewCfg("NatsStreaming"))
		grp.POST("/nias3", api.Factory4NewCfg("Nias3"))
		grp.POST("/benthos", api.Factory4NewCfg("Benthos"))
		grp.POST("/reader", api.Factory4NewCfg("Reader"))
		grp.POST("/align", api.Factory4NewCfg("Align"))
		grp.POST("/textclassifier", api.Factory4NewCfg("TxtClassifier"))
		grp.POST("/level", api.Factory4NewCfg("Level"))
		grp.POST("/weight", api.Factory4NewCfg("Weight"))
		grp.POST("/hub", api.Factory4NewCfg("Hub"))

		// Update Config
		grp.PUT("/natsstreaming", api.UpdateNatsStreamingCfg)
		grp.PUT("/nias3", api.UpdateNias3Cfg)
		grp.PUT("/benthos", api.UpdateBenthosCfg)
		grp.PUT("/reader", api.UpdateReaderCfg)
		grp.PUT("/align", api.UpdateAlignCfg)
		grp.PUT("/textclassifier", api.UpdateTextclassifierCfg)
		grp.PUT("/level", api.UpdateLevelCfg)
		grp.PUT("/weight", api.UpdateWeightCfg)
		grp.PUT("/hub", api.UpdateHubCfg)

		// Delete Config
		grp.DELETE("/natsstreaming", api.DelNatsStreamingCfg)
		grp.DELETE("/nias3", api.DelNias3Cfg)
		grp.DELETE("/benthos", api.DelBenthosCfg)
		grp.DELETE("/reader", api.DelReaderCfg)
		grp.DELETE("/align", api.DelAlignCfg)
		grp.DELETE("/textclassifier", api.DelTextclassifierCfg)
		grp.DELETE("/level", api.DelLevelCfg)
		grp.DELETE("/weight", api.DelWeightCfg)
		grp.DELETE("/hub", api.DelHubCfg)

		if err := e.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil {
			e.Logger.Info("shutting down the server: " + cfg.PageFolder)
		}
	}
}

func main() {

	cfg := config.GetConfig("../config.toml", "./config.toml")
	spew.Dump(cfg)

	go hostPage()()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
