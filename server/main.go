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
		e.POST("/dispense", api.Dispense)
		e.POST("/withdraw", api.Withdraw)
		e.POST("/composite", api.Composite)

		grp := e.Group("/otf-config")
		// grp.Use(middleware.JWT("demoSecret"))

		// Search Config
		grp.GET("/natsstreaming", api.Factory4GetDel("Get", "NatsStreaming"))
		grp.GET("/nias3", api.Factory4GetDel("Get", "Nias3"))
		grp.GET("/benthos", api.Factory4GetDel("Get", "Benthos"))
		grp.GET("/reader", api.Factory4GetDel("Get", "Reader"))
		grp.GET("/align", api.Factory4GetDel("Get", "Align"))
		grp.GET("/textclassifier", api.Factory4GetDel("Get", "TxtClassifier"))
		grp.GET("/level", api.Factory4GetDel("Get", "Level"))
		grp.GET("/weight", api.Factory4GetDel("Get", "Weight"))
		grp.GET("/hub", api.Factory4GetDel("Get", "Hub"))

		// New Config
		grp.POST("/natsstreaming", api.Factory4NewUpdate("New", "NatsStreaming"))
		grp.POST("/nias3", api.Factory4NewUpdate("New", "Nias3"))
		grp.POST("/benthos", api.Factory4NewUpdate("New", "Benthos"))
		grp.POST("/reader", api.Factory4NewUpdate("New", "Reader"))
		grp.POST("/align", api.Factory4NewUpdate("New", "Align"))
		grp.POST("/textclassifier", api.Factory4NewUpdate("New", "TxtClassifier"))
		grp.POST("/level", api.Factory4NewUpdate("New", "Level"))
		grp.POST("/weight", api.Factory4NewUpdate("New", "Weight"))
		grp.POST("/hub", api.Factory4NewUpdate("New", "Hub"))

		// Update Config
		grp.PUT("/natsstreaming", api.Factory4NewUpdate("Update", "NatsStreaming"))
		grp.PUT("/nias3", api.Factory4NewUpdate("Update", "Nias3"))
		grp.PUT("/benthos", api.Factory4NewUpdate("Update", "Benthos"))
		grp.PUT("/reader", api.Factory4NewUpdate("Update", "Reader"))
		grp.PUT("/align", api.Factory4NewUpdate("Update", "Align"))
		grp.PUT("/textclassifier", api.Factory4NewUpdate("Update", "TxtClassifier"))
		grp.PUT("/level", api.Factory4NewUpdate("Update", "Level"))
		grp.PUT("/weight", api.Factory4NewUpdate("Update", "Weight"))
		grp.PUT("/hub", api.Factory4NewUpdate("Update", "Hub"))

		// Delete Config
		grp.DELETE("/natsstreaming", api.Factory4GetDel("Del", "NatsStreaming"))
		grp.DELETE("/nias3", api.Factory4GetDel("Del", "Nias3"))
		grp.DELETE("/benthos", api.Factory4GetDel("Del", "Benthos"))
		grp.DELETE("/reader", api.Factory4GetDel("Del", "Reader"))
		grp.DELETE("/align", api.Factory4GetDel("Del", "Align"))
		grp.DELETE("/textclassifier", api.Factory4GetDel("Del", "TxtClassifier"))
		grp.DELETE("/level", api.Factory4GetDel("Del", "Level"))
		grp.DELETE("/weight", api.Factory4GetDel("Del", "Weight"))
		grp.DELETE("/hub", api.Factory4GetDel("Del", "Hub"))

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
