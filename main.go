package main

func main() {
	logger := InitializeLogging()

	// Doesn't really do anything for now, since server will get interupted
	defer logger.Sync()
	
	StartServer(8080)
}
