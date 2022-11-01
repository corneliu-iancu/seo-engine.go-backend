package http

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// Returns a new router of gin.Engine type.
func NewRouter(logger *zap.Logger) *gin.Engine {
	fmt.Println("[DEBUG] Create new router")

	r := gin.Default()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

	// @todo: reasses if we really need the recovery with zap.
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(logger, true))

	//@todo: use based on env
	// gin.SetMode(gin.DebugMode)

	return r
	// ========================================================
	// OLD Copied code. @todo: remove me
	// ========================================================
	// if config.Debug {
	// gin.SetMode(gin.DebugMode)
	// }
	// engine.HandleMethodNotAllowed = config.HandleMethodNotAllowed
	// engine.Use(buildCorsConf(config.Socket, config.CorsAllowedHeaders))
	// if config.Gzip {
	// 	engine.Use(gzip.Gzip(gzip.DefaultCompression))
	// }
	// ========================================================
}
