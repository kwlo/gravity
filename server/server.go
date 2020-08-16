package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kwlo/gravity/logging"
)

// Server holds the server configurations and context
type Server struct {
	logger logging.Logger
	addr   string
}

// NewServer creates a new struct Server
func NewServer(logger logging.Logger, addr string) *Server {
	return &Server{
		logger: logger,
		addr:   addr,
	}
}

// Start starts up server and listens at port
func (srv *Server) Start() {
	router := chi.NewRouter()

	// Add Middlewares from middlewares.go
	AddMiddlewares(srv, router)

	// Add routes from routes.go
	AddRoutes(srv, router)

	// Add handling static files for UI
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/*", fs)

	srv.logger.Infof("Starting server at: %v", srv.addr)
	http.ListenAndServe(srv.addr, router)
}
