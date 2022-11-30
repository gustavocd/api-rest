# Basic REST API with Go ![](https://github.com/gustavocd/api-rest/workflows/api-rest/badge.svg)

This is a **simple** REST API built with Go.

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
