package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-resty/resty/v2"
)

// Movie represents the relevant fields from OMDb’s JSON.
type Movie struct {
	Title     string `json:"Title"`
	Year      string `json:"Year"`
	Rated     string `json:"Rated"`
	Released  string `json:"Released"`
	Runtime   string `json:"Runtime"`
	Genre     string `json:"Genre"`
	Director  string `json:"Director"`
	Actors    string `json:"Actors"`
	Plot      string `json:"Plot"`
	IMDBRating string `json:"imdbRating"`
	IMDBID    string `json:"imdbID"`
	Response  string `json:"Response"`
	Error     string `json:"Error"`
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <imdb_id> (e.g. tt3896198)")
	}

	imdbID := os.Args[1]
	apiKey := os.Getenv("OMDB_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set OMDB_API_KEY environment variable")
	}

	client := resty.New()
	resp, err := client.R().
		SetQueryParam("apikey", apiKey).
		SetQueryParam("i", imdbID).
		Get("[omdbapi.com](https://www.omdbapi.com/)")

	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	var movie Movie
	if err := json.Unmarshal(resp.Body(), &movie); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	if movie.Response == "False" {
		log.Fatalf("Error: %v", movie.Error)
	}

	fmt.Printf("\n🎬 %s (%s)\n", movie.Title, movie.Year)
	fmt.Printf("⭐ IMDb Rating: %s\n", movie.IMDBRating)
	fmt.Printf("🕒 Runtime: %s\n", movie.Runtime)
	fmt.Printf("🎭 Genre: %s\n", movie.Genre)
	fmt.Printf("🎬 Directed by: %s\n\n", movie.Director)
	fmt.Printf("Plot: %s\n", movie.Plot)
}
