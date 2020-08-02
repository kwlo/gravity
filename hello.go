package main

import (
	"fmt"
	"net/http"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	zap.S().Infof("Served: %s", r.URL.Path[1:])
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.OutputPaths = []string {
		"stdout",
		"zapinfo.log",
	}
	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger)

	defer logger.Sync()
	defer fmt.Printf("ended")

	zap.S().Infof("Starting Server")

	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}
