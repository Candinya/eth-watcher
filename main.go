package main

import (
	"eth-watcher/global"
	"eth-watcher/inits"
	"log"
)

func main() {
	// Initialize everything
	log.Println("Initializing...")

	// Initialize config
	if err := inits.Config(); err != nil {
		log.Fatalln("Failed to load config: ", err)
	}

	// Initialize logger
	if err := inits.Logger(); err != nil {
		log.Fatalln("Failed to load logger: ", err)
	}

	global.Logger.Info("Logger initialized, switch to here.")

	// Initialize redis
	if err := inits.Redis(); err != nil {
		global.Logger.Fatal("Failed to load redis: ", err.Error())
	}

	global.Logger.Debug("Redis initialized, ready to serve.")

	// Initialize jobs
	if err := inits.Jobs(); err != nil {
		global.Logger.Fatal("Failed to start jobs: ", err.Error())
	}

	global.Logger.Info("ETH Watcher started!")

	select {} // Keep process running

}
