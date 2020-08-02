package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	zap.S().Infof("Served: %s", r.URL.Path[1:])
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func StartServer(port int) {
	zap.S().Infof("Starting Server at port %d", port)

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
