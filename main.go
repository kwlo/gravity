package main

import "github.com/kwlo/gravity/pack/level2"
import "github.com/kwlo/gravity/pack"

func main() {
	logger := InitializeLogging()

	// Doesn't really do anything for now, since server will get interupted
	defer logger.Sync()
	
	level2.Test3()
	pack.Test()
	
	StartServer(8080)
}
