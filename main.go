package main

import (
	// printing on console
	"fmt"
	"log"

	// to encode the data for postman
	"encoding/json"
	// for individual unique movie id
	"math/rand"
	// for creating a server in Go
	"net/http"
	// for converting to string
	"strconv"

	"github.com/gorilla/mux"
)

// creating structs as we are not using any database

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

// sending all movies to postman/frontend
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// sending json data to postman
	json.NewEncoder(w).Encode(movies)
}

// deleting a movie with a id
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// taking that params from api request
	// where r is the request that user sends
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// sending a particular movie with a id
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// taking that params from api request
	// where r is the request that user sends
	params := mux.Vars(r)

	// In this we will iterate to all movies and will send to Postman/frontend by converting it to JSON

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	// decoding the json to struct
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// generating random id and then converting it to string
	movie.ID = strconv.Itoa(rand.Intn(10000000))

	// adding to the movies array
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	// Steps for update
	// Set our json content type
	// params
	// loop over the movies , range
	// delete the movie with id  that user sent
	// add a new movie - the movie that we send int the body of postman
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[:index+1]...)

			var movie Movie

			_ = json.NewDecoder(r.Body).Decode(&movie)

			movie.ID = params["id"]

			movies = append(movies, movie)

			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {
	// := declares and defines a variable at the same time
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "43827", Title: "Movie One", Director: &Director{Firstname: "Jhalak", Lastname: "Dashora"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Sharukh", Lastname: "Khan"}})

	// creating routes
	r.HandleFunc("/movies", getMovies).Methods("GET")

	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")

	r.HandleFunc("/movies", createMovie).Methods("POST")

	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Staring server at PORT 4000")
	// listening to server
	log.Fatal(http.ListenAndServe(":4000", r))
}
