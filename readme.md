# Basic REST API with Go

This is a **basic** RESTful Api built using Golang.

## Clone and run the server

```bash
$ git clone https://github.com/gustavocd/api-rest.git && cd api-rest
```

```bash
$ go run api-rest.go
```

## Start making request to the following endpoints with your favorite REST client
    * `/api/notes` -> to get all the notes
    * `/api/notes` -> to create a note
    * `/api/notes/{id}` -> to edit a note
    * `/api/notes/{id}` -> to delete a note
 
 Note: check the source code(models/note.go) to get the correct structure to make the json object :smile: