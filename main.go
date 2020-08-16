package main

import (
	"fmt"
	"os"
	"strconv"

	"net/http"

	"github.com/kwlo/gravity/logging"
	"github.com/kwlo/gravity/server"
)

func getPort() int {
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		return port
	}

	// Returns port 8080 as default
	return 8080
}

func main() {
	logger := logging.NewLogger()

	// Doesn't really do anything for now, since server will get interupted
	defer logger.Sync()

	server := server.NewServer(
		logger,
		fmt.Sprintf("0.0.0.0:%d", getPort()),
		"./static",
		http.ListenAndServe,
	)

	server.Start()
}
