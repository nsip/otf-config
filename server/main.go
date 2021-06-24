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

		// New Config
		grp.POST("/reader", api.NewReaderCfg)
		grp.POST("/align", api.NewAlignCfg)
		grp.POST("/textclassifier", api.NewTextclassifierCfg)
		grp.POST("/level", api.NewLevelCfg)
		grp.POST("/weight", api.NewWeightCfg)
		grp.POST("/hub", api.NewHubCfg)

		// Update Config
		grp.PUT("/reader", api.UpdateReaderCfg)
		grp.PUT("/align", api.UpdateAlignCfg)
		grp.PUT("/textclassifier", api.UpdateTextclassifierCfg)
		grp.PUT("/level", api.UpdateLevelCfg)
		grp.PUT("/weight", api.UpdateWeightCfg)
		grp.PUT("/hub", api.UpdateHubCfg)

		// Delete Config
		grp.DELETE("/reader", api.DelReaderCfg)
		grp.DELETE("/align", api.DelAlignCfg)
		grp.DELETE("/textclassifier", api.DelTextclassifierCfg)
		grp.DELETE("/level", api.DelLevelCfg)
		grp.DELETE("/weight", api.DelWeightCfg)
		grp.DELETE("/hub", api.DelHubCfg)

		// Search Config
		grp.GET("/reader", api.ReaderCfg)
		grp.GET("/align", api.AlignCfg)
		grp.GET("/textclassifier", api.TextclassifierCfg)
		grp.GET("/level", api.LevelCfg)
		grp.GET("/weight", api.WeightCfg)
		grp.GET("/hub", api.HubCfg)

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
