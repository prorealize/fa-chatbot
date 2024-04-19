package main

import (
	"fa-chatbot/api"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const trustedProxy = "0.0.0.0"

func main() {
	_ = godotenv.Load()

	router := api.GetRouters()

	// Set trusted proxies
	trustedProxies := []string{trustedProxy}
	if err := router.SetTrustedProxies(trustedProxies); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// Start the server
	address := serverAddress()
	err := router.Run(address)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

func serverAddress() string {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" {
		host = "0.0.0.0"
	}
	if port == "" {
		port = "8080"
	}
	address := fmt.Sprintf("%s:%s", host, port)
	return address
}
