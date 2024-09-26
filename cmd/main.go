package main

import (
	"log"
	"net/http"
	"re-partners/internal/app"
	"re-partners/internal/app/handler"
	"re-partners/internal/app/service"
	"re-partners/internal/config"
)

func main() {
	fulfillment := handler.New(service.New())
	router := app.NewRouter(fulfillment)

	port := ":" + config.GetEnv().Server.Port
	if port == "" {
		port = ":8080"
	}
	log.Println("server is starting...")
	log.Fatal(http.ListenAndServe(port, router))
}
