package main

import (
	"os"
	"strconv"
)

func getPort() int {
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		return port
	}

	// Returns port 8080 as default
	return 8080
}

func main() {
	logger := NewLogger()

	// Doesn't really do anything for now, since server will get interupted
	defer logger.Sync()

	server := &Server{logger: logger}
	server.Start(getPort())
}
