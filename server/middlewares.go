package server

import (
	"net/http"

	"github.com/go-chi/chi"
)

// AddMiddlewares adds middleware to the router
func AddMiddlewares(srv *Server, router *chi.Mux) {
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: blank middleware for now
			next.ServeHTTP(w, r)
		})
	})
}
