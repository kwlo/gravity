package main

import (
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

// HelloServer Sample route. To be deleted
func HelloServer(w http.ResponseWriter, r *http.Request) {
	zap.S().Infof("Served: %s", r.URL.Path[1:])
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

// StartServer Start up server and listen at port
func StartServer(port int) {
	zap.S().Infof("Starting Server at port %d", port)

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
