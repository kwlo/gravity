package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// AddRoutes adds the http routes to the router
func AddRoutes(srv *Server, router *chi.Mux) {
	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})

	router.Get("/simulations/{simulationID}", func(w http.ResponseWriter, r *http.Request) {
		simulationID := chi.URLParam(r, "simulationID")
		srv.logger.Infof("Serving simulation ID: %v", simulationID)
		fmt.Fprintf(w, "ID: %v\n", simulationID)
	})
}
