package models

// Note holds the data of notes created by the users.
type Note struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedOn   string `json:"createdOn"`
}
