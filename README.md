

# CRUD API with Golang

This is a simple CRUD (Create, Read, Update, Delete) API for managing movies using Golang. It doesn't use a database, but rather manipulates data stored in memory using Go structs and slices.

## Features

- **GET /movies**: Retrieve all movies.
- **GET /movies/{id}**: Retrieve a specific movie by ID.
- **POST /movies**: Create a new movie.
- **PUT /movies/{id}**: Update a specific movie by ID.
- **DELETE /movies/{id}**: Delete a specific movie by ID.

## Getting Started

### Prerequisites

- Go 1.16+
- Gorilla Mux

### Installing Gorilla Mux

To install the necessary dependencies, run:

```bash
go get -u github.com/gorilla/mux
```

### Running the API

1. Clone the repository or copy the code into a Go project.
2. Run the application:

```bash
go run main.go
```

3. The server will start on port `8000`. Access the API at `http://localhost:8000`.

### API Endpoints

| Method | Endpoint          | Description               |
|--------|-------------------|---------------------------|
| GET    | `/movies`          | Get all movies            |
| GET    | `/movies/{id}`     | Get a movie by ID         |
| POST   | `/movies`          | Create a new movie        |
| PUT    | `/movies/{id}`     | Update a movie by ID      |
| DELETE | `/movies/{id}`     | Delete a movie by ID      |

### Example Requests

- **Create a Movie (POST)**

```bash
curl -X POST http://localhost:8000/movies \
-H "Content-Type: application/json" \
-d '{"Isbn": "12345", "Title": "New Movie", "Director": {"Firstname": "John", "Lastname": "Doe"}}'
```

- **Get All Movies (GET)**

```bash
curl -X GET http://localhost:8000/movies
```

### Note

This API uses in-memory data storage, meaning all data will be lost once the server is stopped.
