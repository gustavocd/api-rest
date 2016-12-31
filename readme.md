# Basic REST API with Go

This is a **basic** RESTful Api built using the go programming language.

## Usage
1. Download or clone the repository
2. `cd <folder-name>` and then press enter
3. Type the command `go build api-rest.go`
4. Execute the binaries `api-rest.exe` on Windows and `./api-rest` on Linux and Mac
5. Star making request to the following end points with your favourite tool
    * `/api/notes` -> to get all the notes
    * `/api/notes` -> to create a note
    * `/api/notes/{id}` -> to edit a note
    * `/api/notes/{id}` -> to delete a note
 
 Note: check the source code(models/note.go) to get the correct structure to make the json object :smile