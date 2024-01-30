package main

import (
	"log"
	"net/http"

	"github.com/gustavocd/api-rest/routes"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/notes", routes.GetNoteHandler)
	mux.HandleFunc("POST /api/notes", routes.PostNoteHandler)
	mux.HandleFunc("PUT /api/notes/{id}", routes.PutNoteHandler)
	mux.HandleFunc("DELETE /api/notes/{id}", routes.DeleteNoteHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("Server running at http://localhost:8080")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
