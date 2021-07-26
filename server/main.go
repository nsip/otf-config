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
		e.POST("/composite", api.Composite)

		grp := e.Group("/otf-config")
		// grp.Use(middleware.JWT("demoSecret"))

		// Search Config
		grp.GET("/", api.Factory4GetDel("Get"))

		// New Config
		grp.POST("/", api.Factory4NewUpdate("New"))

		// Update Config
		grp.PUT("/", api.Factory4NewUpdate("Update"))

		// Delete Config
		grp.DELETE("/", api.Factory4GetDel("Del"))

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
