package routes

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetNotes(t *testing.T) {
	t.Run("should return null when there are no notes", func(t *testing.T) {
		r, err := http.NewRequest("GET", "/api/notes", nil)
		if err != nil {
			t.Fatal(err)
		}
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(GetNoteHandler)
		handler.ServeHTTP(w, r)
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %d want %d", status, http.StatusOK)
		}
		expected := `null`
		if w.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %s want %s", w.Body.String(), expected)
		}
	})
}

func TestPostNote(t *testing.T) {
	t.Run("should create a note when the sent data is correct", func(t *testing.T) {
		var note = []byte(`{"title": "Primera nota","description": "Esta es la descripción de mi primera nota"}`)
		r, err := http.NewRequest("POST", "/api/notes", bytes.NewBuffer(note))
		if err != nil {
			t.Fatal(err)
		}
		r.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(PostNoteHandler)
		handler.ServeHTTP(w, r)
		if status := w.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %d want %d", status, http.StatusCreated)
		}
		expected := fmt.Sprintf(`{"id":"1","title":"Primera nota","description":"Esta es la descripción de mi primera nota","createdOn":"%v"}`, time.Now().Format(time.RFC3339))
		if w.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %s want %s", w.Body.String(), expected)
		}
	})
}
