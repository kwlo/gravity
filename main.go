package main

import "github.com/kwlo/gravity/physics"

func main() {
	logger := InitializeLogging()

	// Doesn't really do anything for now, since server will get interupted
	defer logger.Sync()

	physics.ConvertDistance(10, physics.Yards, physics.Inches)

	StartServer(8080)
}
