package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Returns a new router of gin.Engine type.
func NewRouter() *gin.Engine {
	fmt.Println("[DEBUG] Create new router")

	engine := gin.Default()
	gin.SetMode(gin.DebugMode)

	return engine
	// if config.Debug {
	// gin.SetMode(gin.DebugMode)
	// }
	// engine.HandleMethodNotAllowed = config.HandleMethodNotAllowed
	// engine.Use(buildCorsConf(config.Socket, config.CorsAllowedHeaders))
	// if config.Gzip {
	// 	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	// }
	//
}
