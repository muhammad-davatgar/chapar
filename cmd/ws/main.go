package ws

import (
	messangerservice "chapar/internals/core/services/messanger"
	"chapar/internals/handlers/gorilla"
	"log"
	"net/http"
	"time"
)

func main() {
	msgService := messangerservice.NewMessangerService()

	gorillaService := gorilla.NewGorillaService

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	server := &http.Server{
		Addr:              *addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
