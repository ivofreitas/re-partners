package main

import (
	"log"
	"net/http"
	"os"
	"re-partners/internal/app"
	"re-partners/internal/app/handler"
	"re-partners/internal/app/service"
)

func main() {
	fulfillment := handler.New(service.New())
	router := app.NewRouter(fulfillment)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = ":8080"
	}
	log.Println("server is starting...")
	log.Fatal(http.ListenAndServe(port, router))
}
