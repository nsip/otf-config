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

	cfg := config.OtfCfg

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
		grp.GET("/natsstreaming", api.NatsStreamingCfg)
		grp.GET("/nias3", api.Nias3Cfg)
		grp.GET("/benthos", api.BenthosCfg)
		grp.GET("/reader", api.ReaderCfg)
		grp.GET("/align", api.AlignCfg)
		grp.GET("/textclassifier", api.TextclassifierCfg)
		grp.GET("/level", api.LevelCfg)
		grp.GET("/weight", api.WeightCfg)
		grp.GET("/hub", api.HubCfg)

		// New Config
		grp.POST("/natsstreaming", api.NewNatsStreamingCfg)
		grp.POST("/nias3", api.NewNias3Cfg)
		grp.POST("/benthos", api.NewBenthosCfg)
		grp.POST("/reader", api.NewReaderCfg)
		grp.POST("/align", api.NewAlignCfg)
		grp.POST("/textclassifier", api.NewTextclassifierCfg)
		grp.POST("/level", api.NewLevelCfg)
		grp.POST("/weight", api.NewWeightCfg)
		grp.POST("/hub", api.NewHubCfg)

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

	cfg := config.OtfCfg
	spew.Dump(cfg)

	go hostPage()()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
}
