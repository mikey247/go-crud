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

type Movie struct {
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func getMovie(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(req)

	for _,item:= range movies{
		if item.ID==params["id"]{
			json.NewEncoder(res).Encode(item)
			break
		}
	}
}

func createMovie(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)
	json.NewEncoder(res).Encode(movie)
}

func updateMovie(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "application/json")
	//params
	params:= mux.Vars(req)
	// loop over movies 
	for index,item:= range movies{
		if item.ID==params["id"]{
			// delete movie with id sent 
			movies = append(movies[:index], movies[index+1:]...)
			//add a new movie
			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(100000000))
			movies = append(movies, movie)
			json.NewEncoder(res).Encode(movie)
			return
		}
	}
}

func deleteMovie(res http.ResponseWriter, req *http.Request)  {
	res.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(req)
	
	for index,item:= range movies{
		if item.ID==params["id"]{
			movies=append(movies[:index], movies[index+1:]... )
			break
		}
	}
	json.NewEncoder(res).Encode(movies)
}

func main()  {
	r:= mux.NewRouter()

	movies = append(movies, Movie{ID:"1", Isbn: "ISbn14", Title: "Movie One", Director: &Director{Firstname: "Mikey", Lastname: "Ume"} })
	movies = append(movies, Movie{ID:"2", Isbn: "ISbn24", Title: "Movie Two", Director: &Director{Firstname: "Emma", Lastname: "Ume"} })
	movies = append(movies, Movie{ID:"3", Isbn: "ISbn34", Title: "Movie Three", Director: &Director{Firstname: "Dan", Lastname: "Ume"} })
	movies = append(movies, Movie{ID:"1", Isbn: "ISbn44", Title: "Movie Four", Director: &Director{Firstname: "Sam", Lastname: "Ume"} })

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")   
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting Server at port 8000")
	log.Fatal(http.ListenAndServe(":8000",r))
}