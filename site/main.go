package main

import (
	"net/http"
)

type Spotify_Donnee struct {
	Q1 string
	Q2 string
	Q3 string
	Q4 string
	Q5 string
	Q6 string
	Q7 string
}

func main() {
	http.HandleFunc("/", HandleHome)
	API_Q1 := "https://api.spotify.com/v1/search?q=Jul&type=artist&limit=3&market=FR"
	API_Q2 := "https://api.spotify.com/v1/artists/3IW7ScrzXmPvZhB27hmfgy/albums"
	API_Q3 := "https://api.spotify.com/v1/albums/1WYGwCvsfFrr7kuQcfNnJr/tracks"
	API_Q4 := "https://api.spotify.com/v1/search?q=Dans+ma+parano%C3%AFa&type=track"
	API_Q5 := "https://api.spotify.com/v1/search?q=SCH&type=artist&market=FR"
	API_Q6 := "https://api.spotify.com/v1/search?q=Bolide+allemand&type=track"
	API_Q7 := "https://api.spotify.com/v1/search?q=%E2%80%9CLes+actus+du+jour+-+Hugo+d%C3%A9crypte&type=track"
	response_1, err := http.Get(API_Q1)
	if err != nil {
		fmt.Println("Erreur", err)
		return
	}
	defer response_1.Body.Close()

	if response_1.StatusCode != http.StatusOK {
		fmt.Println("Réponse Q1", response_1.Status)
		return
	}



	

	http.ListenAndServe("localhost:8080", nil)
}

