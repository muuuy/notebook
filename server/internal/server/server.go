package server

import (
	"fmt"
	"notepad/internal/config"
)

func Start(cfg *config.Config) error {
	fmt.Printf("Starting server on port %s...\n", cfg.ServerPort)
	// Add server startup logic here (e.g., using net/http)
	return nil // Return an error if something goes wrong
}
