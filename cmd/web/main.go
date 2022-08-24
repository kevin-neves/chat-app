package main

import (
	"log"
	"net/http"

	"github.com/kevin-neves/chat-app/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("Serving chat server on port 8080")
	err := http.ListenAndServe(":8080", mux)

	if err != nil {
		log.Println(err)
	}
}
