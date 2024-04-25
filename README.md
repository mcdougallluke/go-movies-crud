# Movie CRUD API in Golang

This project is a simple CRUD (Create, Read, Update, Delete) API for managing movie information, built using Golang and the Gorilla Mux router. The API allows clients to interact with movie data stored in an in-memory slice.

## Features

- List all movies
- Retrieve a single movie by ID
- Add a new movie
- Update an existing movie
- Delete a movie

## Technologies Used

- Golang
- Gorilla Mux
- JSON encoding and decoding

## Setup

To run this project, install Go and set up your Go workspace. Then clone this repository into your Go workspace:

```bash
git clone https://github.com/mcdougallluke/go-movies-crud
cd go-movies-crud
```

## Running the Server
To start the server, run the following command from the root of your project:

```bash
go build
go run main.go
```
This will start the server on http://localhost:8000 with a confirmation message, "Server started on port 8000".

## API Endpoints

The following endpoints are available:

### GET /movies
- Description: Get a list of all movies.
- Response: JSON array of movie objects.

### GET /movies/{id}
- Description: Get a single movie by ID.
- Response: JSON object of a movie.

### POST /movies
- Description: Add a new movie.
- Request body: JSON object of a movie without ID (ID is generated automatically).
- Response: JSON array of all movies, including the newly added one.

### PUT /movies/{id}
- Description: Update an existing movie.
- Request body: JSON object with updated data for the movie.
- Response: JSON array of all movies, including the updated one.

### DELETE /movies/{id}
- Description: Delete a movie by ID.
- Response: JSON array of remaining movies after deletion.

## Example Usage
You can use tools like curl or Postman to interact with the API. Here are some curl commands to get you started:

```bash
# Get all movies
curl http://localhost:8000/movies

# Get a movie by ID
curl http://localhost:8000/movies/1

# Add a new movie
curl -X POST -d '{"isbn":"789012", "title":"Movie Three", "director":{"firstname":"Alice", "lastname":"Johnson"}}' http://localhost:8000/movies

# Update a movie
curl -X PUT -d '{"isbn":"789012", "title":"Updated Movie Three", "director":{"firstname":"Alice", "lastname":"Johnson"}}' http://localhost:8000/movies/3

# Delete a movie
curl -X DELETE http://localhost:8000/movies/
```