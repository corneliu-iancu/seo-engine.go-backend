package http

import (
	"fmt"
	"github.com/gin-contrib/cors"
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

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "PUT", "PATCH", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

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
