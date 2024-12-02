package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"notepad/internal/config"
	"notepad/internal/server"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

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

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	err = http.ListenAndServe("127.0.0.1:3000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server  closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
