package main

import (
	"fmt"
	"os"
	"strconv"

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

	server := &server.Server{
		Addr:   fmt.Sprintf("0.0.0.0:%d", getPort()),
		Logger: logger,
	}
	server.Start()
}
