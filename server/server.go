package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/kwlo/gravity/logging"
)

// Server holds the server configurations and context
type Server struct {
	Logger logging.Logger
}

// Start starts up server and listens at port
func (srv *Server) Start(port int) {
	r := chi.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	r.Handle("/", fs)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	r.Get("/simulations/{simulationID}", func(w http.ResponseWriter, r *http.Request) {
		simulationID := chi.URLParam(r, "simulationID")
		srv.Logger.Infof("Serving simulation ID: %v", simulationID)
		fmt.Fprintf(w, "ID: %v\n", simulationID)
	})

	addr := fmt.Sprintf(":%d", port)
	srv.Logger.Infof("Starting server at port: ", addr)
	http.ListenAndServe(addr, r)
}
