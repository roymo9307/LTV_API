package main

import (
	"net/http"

	"fmt"

	"goji.io/pat"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getAllsongs(w http.ResponseWriter, r *http.Request) {
	z := allSongs()
	w.Write(z)
}

func getByArtist(w http.ResponseWriter, r *http.Request) {
	var artist = pat.Param(r, "artist")
	response := byArtist(artist)
	w.Write(response)
}

func getbysongID(w http.ResponseWriter, r *http.Request) {
	var songID = pat.Param(r, "id")
	response := bySongID(songID)
	w.Write(response)
}

func getBySong(w http.ResponseWriter, r *http.Request) {
	var song = pat.Param(r, "song")
	response := bySong(song)
	w.Write(response)
}

func getByGenre(w http.ResponseWriter, r *http.Request) {
	var genre = pat.Param(r, "genre")
	response := byGenre(genre)
	w.Write(response)
}

func getAllGenres(w http.ResponseWriter, r *http.Request) {
	response := allGenres()
	w.Write(response)
}

func getByGenereID(w http.ResponseWriter, r *http.Request) {
	var genreID = pat.Param(r, "id")
	response := byGenreID(genreID)
	w.Write(response)
}

func getGenresSummary(w http.ResponseWriter, r *http.Request) {
	response := genresSummary()
	w.Write(response)
}

func getByLength(w http.ResponseWriter, r *http.Request) {
	min := pat.Param(r, "min")
	max := pat.Param(r, "max")
	response := byLength(min, max)
	w.Write(response)
}
