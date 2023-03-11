package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func deleteMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(writer).Encode(movies)
}

func updateMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(writer).Encode(movie)
			return
		}
	}
}

func createMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
}

func getMovie(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	params := mux.Vars(request)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
}

func getMovies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = append(
		movies,
		Movie{
			ID:    "1",
			ISBN:  "438227",
			Title: "Movie One",
			Director: &Director{
				FirstName: "John",
				LastName:  "Adams",
			},
		},
	)
	movies = append(
		movies,
		Movie{
			ID:    "2",
			ISBN:  "438228",
			Title: "Movie Two",
			Director: &Director{
				FirstName: "Christopher",
				LastName:  "Nolan",
			},
		},
	)
	movies = append(
		movies,
		Movie{
			ID:    "3",
			ISBN:  "438229",
			Title: "Movie Three",
			Director: &Director{
				FirstName: "Mayim",
				LastName:  "Bialik",
			},
		},
	)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/id", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/id", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/id", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at Port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

