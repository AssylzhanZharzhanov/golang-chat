package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Server started...")
	r := mux.NewRouter()
	room := newRoom()
	go room.run()

	r.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(room, w, r)
	})

	log.Fatal(http.ListenAndServe(":5000", r))
}
