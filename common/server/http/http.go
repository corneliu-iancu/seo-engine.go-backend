package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
func (s *Server) Start(ctx context.Context) {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		Handler:      s.router, // gin router
		ReadTimeout:  s.timeout,
		WriteTimeout: s.timeout,
	}

	fmt.Println("[DEBUG] 🚀 HTTP server started on port 9000.")

	// make use of go channels.
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("[FATAL] Server stoped due to errors")
		os.Exit(1)
	}
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
