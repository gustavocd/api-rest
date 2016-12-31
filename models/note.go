package models

import "time"

//Note holds the data of notes created by the users.
type Note struct {
	Title string `json:"title"`
	Description string `json:"description"`
	CreatedOn time.Time `json:"createdon"`
}
