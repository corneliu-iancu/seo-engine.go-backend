package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// HTTP Server struct.
type Server struct {
	port    int32
	timeout time.Duration
	router  *gin.Engine
	// log     Logger
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
			log.Fatalf("[DEBUG] Listen:%+s\n", err)
		}
	}()

	fmt.Println("[DEBUG] 🚀 HTTP server started on port 9000.")

	<-ctx.Done()

	log.Printf("[DEBUG] 🚦 Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = server.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("[DEBUG] 🔴 Server Shutdown Failed:%+s", err)
	}

	log.Printf("[DEBUG] ✅ Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

// () config *Config, log Logger, monitor Monitor
func New(routes func(*gin.Engine)) *Server {
	// Default configuration.
	port := 9000
	timeout := 10 * time.Second // default timeout - 10 000 milliseconds (10s)

	// Creates new gin router.
	router := NewRouter()

	// Apply custom routes handler to the gin router.
	routes(router)

	// Returns a Server object.
	return &Server{
		port:    int32(port),
		timeout: timeout,
		router:  router,
	}
}

// Initialize the HTTP Server.
func Init(routes func(*gin.Engine)) *Server {
	// Init function.
	// @Todo: Add server configuration.
	return New(routes)
}
