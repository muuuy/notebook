package main

import (
	"log"
	"notepad/internal/config"
	"notepad/internal/server"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Start the server
	if err := server.Start(cfg); err != nil { // Assuming server.Start returns an error
		log.Fatalf("Error starting server: %v", err)
	}
}
