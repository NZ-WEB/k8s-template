package main

import (
	"k8s-template/internal/handlers"
	"k8s-template/pkg/version"
	"log"
	"net/http"
	"os"
)

// How to try it: PORT=8000 go run main.go
func main() {
	log.Printf(
		"Starting the service...\ncommit: %s, build time: %s, release: %s",
		version.Commit, version.BuildTime, version.Release,
	)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port is not set.")
	}

	r := handlers.Router(version.BuildTime, version.Commit, version.Release)
	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
