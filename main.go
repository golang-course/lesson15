package main

import (
	"lesson15/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var r = mux.NewRouter()
	r.HandleFunc("/", handlers.IndexHandler)
	r.HandleFunc("/unknown", handlers.UnknownHandler)
	r.HandleFunc("/about", handlers.AboutHandler)
	r.HandleFunc("/contacts", handlers.ContactsHandler)
	r.HandleFunc("/api/phones", handlers.AllPhonesHandler)
	r.HandleFunc("/api/phones/{id:[0-9]+}", handlers.OnePhoneHandler)

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Println(err)
		return
	}
}
