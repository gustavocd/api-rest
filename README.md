# Basic REST API with Go [![Build Status](https://travis-ci.com/gustavocd/api-rest.svg?branch=master)](https://travis-ci.com/gustavocd/api-rest)

This is a **basic** RESTful Api built using Golang.

## Clone and run the server

```bash
git clone https://github.com/gustavocd/api-rest.git && cd api-rest
```

```bash
make run
```

## Available endpoints

| Endpoint          | Description       | HTTP method |
|-------------------|-------------------|-------------|
| `/api/notes`      | Get all the notes | `GET`       |
| `/api/notes`      | Create a note     | `POST`      |
| `/api/notes/{id}` | Edit a note       | `PUT`       |
| `/api/notes/{id}` | Delete a note     | `DELETE`    |
 
 Note: check the source code(models/note.go) to get the correct structure to make the json object :smile:
