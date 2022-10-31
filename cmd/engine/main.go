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
	"log"
	"os"
	"os/signal"

	"github.com/corneliu-iancu/seo-engine.go-backend/common/server/http"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/factory"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/handler"
)

// main entry point to the go app.
// @todo: Should be moved to superbet Task implementation, aloing with monitoring and logging.
func main() {
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
	httpServer := http.Init(handler.GinHandler(httpControllers))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		oscall := <-c
		log.Printf("[DEBUG]️️ ⚡️ System call:%+v", oscall)
		cancel()
	}()

	if err := httpServer.Start(ctx); err != nil {
		log.Printf("[ERROR] 🔴 Failed to serve:+%v\n", err)
	}
}
