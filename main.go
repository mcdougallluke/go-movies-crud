package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Movie Struct
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director Struct
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// return the movies array as json
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// get the id from the request params
	params := mux.Vars(r)

	// loop through the movies array and delete the movie with the id
	for index, item := range movies {
		if item.ID == params["id"] {
			// delete the movie from the movies array
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// return the movies array as json
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// get the id from the request params
	params := mux.Vars(r)

	// loop through the movies array and return the movie with the id
	for _, item := range movies {
		if item.ID == params["id"] {
			// return the movie as json
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// create a new movie object
	var movie Movie
	// decode the request body to the movie object
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// set the movie id
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	// append the movie to the movies array
	movies = append(movies, movie)
	// return the movies array as json
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// set the header to application/json
	w.Header().Set("Content-Type", "application/json")
	// get the id from the request params
	params := mux.Vars(r)

	// loop through the movies array and update the movie with the id
	for index, item := range movies {
		if item.ID == params["id"] {
			// delete the movie from the movies array
			movies = append(movies[:index], movies[index+1:]...)
			// create a new movie object
			var movie Movie
			// decode the request body to the movie object
			_ = json.NewDecoder(r.Body).Decode(&movie)
			// set the movie id
			movie.ID = params["id"]
			// append the movie to the movies array
			movies = append(movies, movie)
			// return the movies array as json
			json.NewEncoder(w).Encode(movies)
			return
		}
	}
	// return the movies array as json
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	// Mock Data
	movies = append(movies, Movie{ID: "1", Isbn: "123456", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "654321", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server started on port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
