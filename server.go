package main

import (
	"fmt"
	"net/http"
)

// Server holds the server configurations and context
type Server struct {
	logger Logger
}

// Route Sample route. To be deleted
func (srv *Server) Route(w http.ResponseWriter, r *http.Request) {
	srv.logger.Infof("Served: %s", r.URL.Path[1:])
	fmt.Fprintf(w, "pong")
}

// Start starts up server and listens at port
func (srv *Server) Start(port int) {
	srv.logger.Infof("Starting Server at port %d", port)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/ping", srv.Route)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
