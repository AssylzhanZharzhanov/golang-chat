package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	room := newRoom()

	r.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		serveWs(room, w, r)
	})

	log.Fatal(http.ListenAndServe(":5000", r))
}
