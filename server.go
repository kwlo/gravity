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
func (v *Server) Route(w http.ResponseWriter, r *http.Request) {
	v.logger.Infof("Served: %s", r.URL.Path[1:])
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// Start starts up server and listens at port
func (v *Server) Start(port int) {
	v.logger.Infof("Starting Server at port %d", port)

	http.HandleFunc("/", v.Route)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
