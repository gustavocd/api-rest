package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gustavocd/api-rest/routes"
)

// Entry point of program
func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/api/notes", routes.GetNoteHandler).Methods("GET")
	r.HandleFunc("/api/notes", routes.PostNoteHandler).Methods("POST")
	r.HandleFunc("/api/notes/{id}", routes.PutNoteHandler).Methods("PUT")
	r.HandleFunc("/api/notes/{id}", routes.DeleteNoteHandler).Methods("DELETE")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("Server running at http://localhost:8080")
	server.ListenAndServe()
}
