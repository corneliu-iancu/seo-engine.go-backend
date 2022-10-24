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

	"github.com/corneliu-iancu/seo-engine.go-backend/common/server/http"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/factory"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/handler"
)

// main entry point to the go app.
// @todo: Should be moved to superbet Task implementation, aloing with monitoring and logging.
func main() {

	// application instance.
	services := factory.NewApplication()

	// create rules table.
	// err := services.CreateRulesTable()
	// if err != nil {
	// 	fmt.Println("[ERROR] ", err)
	// }
	// os.Exit(0)

	// http handlers
	httpControllers := handler.NewHttpControllers(services)

	// http server with our custom route config
	httpServer := http.Init(handler.GinHandler(httpControllers))

	// Starts the actual HTTP server.
	ctx, _ := context.WithCancel(context.Background())
	httpServer.Start(ctx)
}
