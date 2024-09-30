package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"golang.org/x/exp/rand"
)

/*
I will not be using database because I do not want it to get complex right now!
I will be using structs and slices to manipulate the data inside go only!
*/

// structs in go are like objects in javascript, it will be having key value pairs!

// we have two structs initially : movies and directors and they are going to be related to each other in the sense that every movie has a director

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json :"isbn"`
	Title    string    `json : "title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json : "lastname"`
}

// defining a variable movie of the type : slice Movie because we will be using movies very often

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["ID"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
}

func updateMovie() {

}

func main() {
	r := mux.NewRouter()

	/*
	   when we will go to postman and hit the server ,when we will hit the api /movies and we want movies
	   in the beginning , there won't be any movies and we don't want that !
	*/

	movies = append(movies, Movie{ID: "1", Isbn: "43434", Title: "Movie one", Director: &Director{Firstname: "ek", Lastname: "Villain"}})
	movies = append(movies, Movie{ID: "2", Isbn: "23232", Title: "Movie ntwo", Director: &Director{Firstname: "SACHIN", Lastname: "Tendulkar"}})

	r.handleFunc("/movies", getMovies).Methods("GET")
	r.handleFunc("/movies/{id}", getMovie).Methods("GET")
	r.handleFunc("/movies", createMovie).Methods("POST")
	r.handleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.handleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("startign server at port : 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
