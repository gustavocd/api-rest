package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gustavocd/api-rest/models"
)

// Store for the Notes collection
var noteStore = make(map[string]models.Note)

// Variable to generate key for the collection
var id int

// PostNoteHandler HTTP Post - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note models.Note
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedOn = time.Now().Format(time.RFC3339)
	id++
	k := strconv.Itoa(id)
	note.ID = k
	noteStore[k] = note

	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write(j)
	if err != nil {
		panic(err)
	}
}

// GetNoteHandler HTTP Get - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	var notes []models.Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(j)
	if err != nil {
		panic(err)
	}
}

// PutNoteHandler HTTP Put - /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	k := r.PathValue("id")
	var noteToUpd models.Note
	// Decode the incoming Note JSON
	err = json.NewDecoder(r.Body).Decode(&noteToUpd)
	if err != nil {
		panic(err)
	}

	note, ok := noteStore[k]
	if !ok {
		log.Printf("Could not find key of Note %s to delete", k)
		return
	}

	noteToUpd.CreatedOn = note.CreatedOn
	noteToUpd.ID = note.ID
	// Delete existing item and add the updated item
	delete(noteStore, k)
	noteStore[k] = noteToUpd
	w.WriteHeader(http.StatusNoContent)
}

// DeleteNoteHandler HTTP Delete - /api/notes/{id}
func DeleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	k := r.PathValue("id")
	//Remove from store
	if _, ok := noteStore[k]; ok {
		//delete the existing item
		delete(noteStore, k)
	} else {
		log.Printf("Could not find key of Note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}
