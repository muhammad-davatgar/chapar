package main

import (
	messangerservice "chapar/internals/core/services/messanger"
	"chapar/internals/handlers/gorilla"
	"log"
	"net/http"
	"time"
)

func main() {
	msgService := messangerservice.NewMessangerService()

	gorillaService := gorilla.NewGorillaService(msgService)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		gorillaService.ServeWs(w, r)
	})

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
