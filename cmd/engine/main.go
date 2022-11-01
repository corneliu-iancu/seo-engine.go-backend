// Package for the SEO engine.
// adding new SEO rules with custom metadata
// based on a given URI returns the proper metadata information
//
// app steps:
//
//	> initialze app services
//	> initialize http server.
//
// @todo:
//
//	implement logging
//	implement monitoring
package main

import (
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"

	"github.com/corneliu-iancu/seo-engine.go-backend/common/server/http"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/factory"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/handler"
)

// main entry point to the go app.
func main() {
	// @todo: based on dev environment.
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// application instance.
	services := factory.NewApplication()

	// =============================================
	// @todo: move me to other me part of the app.
	// @todo: handle resource not found error.
	// =============================================
	// create rules table.
	// =============================================
	// err := services.CreateRulesTable()
	// if err != nil {
	//	log.Printf("[ERROR] %+v", err)
	// }
	// os.Exit(0)
	// =============================================

	// http handlers
	httpControllers := handler.NewHttpControllers(services)

	// http server with our custom route config
	httpServer := http.Init(handler.GinHandler(httpControllers), logger)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		_ = <-c
		oscall := <-c
		logger.Info("[DEBUG]️️ ⚡️ System call:%+v", zap.String("oscall", oscall.String()))
		cancel()
	}()

	if err := httpServer.Start(ctx); err != nil {
		logger.Error("[ERROR] 🔴 Failed to serve:+%v\n", zap.Error(err))
	}
}
