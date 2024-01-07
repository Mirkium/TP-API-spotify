package main

import (
	"net/http"
)

type Spotify_Donnee struct {
	Artist string
	Track  string
}

func main() {
	http.HandleFunc("/", HandleHome)

	http.ListenAndServe("localhost:8080", nil)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	API := "https://accounts.spotify.com/api/token"

	spotifyData := Spotify_Donnee{
		Artist: "Nom de l'artiste",
		Track:  "Nom du morceau",
	}
}
