package server

import (
	"fmt"
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
	r := chi.NewRouter()

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	r.Get("/simulations/{simulationID}", func(w http.ResponseWriter, r *http.Request) {
		simulationID := chi.URLParam(r, "simulationID")
		srv.logger.Infof("Serving simulation ID: %v", simulationID)
		fmt.Fprintf(w, "ID: %v\n", simulationID)
	})

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/*", fs)

	srv.logger.Infof("Starting server at: %v", srv.addr)
	http.ListenAndServe(srv.addr, r)
}
