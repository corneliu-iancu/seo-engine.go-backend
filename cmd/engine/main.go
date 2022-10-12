// todo: Main package for the SEO engine.
// - initialize http server.
package main

import (
	"context"
	"github.com/corneliu-iancu/seo-engine.go-backend/common/server/http"
	"github.com/corneliu-iancu/seo-engine.go-backend/internal/handler"
)

func main() {
	httpServer := http.Init(handler.GinHandler())

	// @TODO: Should be moved to superbet Task implementation, aloing with monitoring and logging.

	// Starts the actual HTTP server.
	ctx, _ := context.WithCancel(context.Background())
	httpServer.Start(ctx)
}
