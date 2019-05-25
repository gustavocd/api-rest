package models

// Note holds the data of notes created by the users.
type Note struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedOn   string `json:"createdon"`
}
