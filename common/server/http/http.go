package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// HTTP Server struct.
type Server struct {
	port    int32
	timeout time.Duration
	router  *gin.Engine
	log     *zap.Logger
	// monitor Monitor
}

// Start starts listening for requests and serving responses.
func (s *Server) Start(ctx context.Context) (err error) {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.router, // gin router
		ReadTimeout:  s.timeout,
		WriteTimeout: s.timeout,
	}

	go func() {
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.log.Fatal("[DEBUG] Listen:%+s\n", zap.Error(err))
		}
	}()

	s.log.Debug("🚀 HTTP server started on port 9000.")

	<-ctx.Done()

	s.log.Debug("🚦 Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = server.Shutdown(ctxShutDown); err != nil {
		s.log.Fatal("[DEBUG] 🔴 Server Shutdown Failed:%+s", zap.Error(err))
	}

	s.log.Debug("✅ Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

// () config *Config, log Logger, monitor Monitor
func New(routes func(*gin.Engine), logger *zap.Logger) *Server {
	// Default configuration.
	port := 9000
	timeout := 10 * time.Second // default timeout - 10 000 milliseconds (10s)

	// Creates new gin router.
	router := NewRouter(logger)

	// Apply custom routes handler to the gin router.
	routes(router)

	// Returns a Server object.
	return &Server{
		port:    int32(port),
		timeout: timeout,
		router:  router,
		log:     logger,
	}
}

// Initialize the HTTP Server.
func Init(routes func(*gin.Engine), logger *zap.Logger) *Server {
	// Init function.
	// @Todo: Add server configuration.
	return New(routes, logger)
}
