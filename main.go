package main

/*
This page is in charge communication
between the app and web (view) using JSON
*/

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), homePage)

	mux.HandleFunc(pat.Get("/songs"), getAllsongs)
	mux.HandleFunc(pat.Get("/songs/id/:id"), getbysongID)

	mux.HandleFunc(pat.Get("/songs/artist/:artist"), getByArtist)
	mux.HandleFunc(pat.Get("/songs/song/:song"), getBySong)
	mux.HandleFunc(pat.Get("/songs/genre/:genre"), getByGenre)

	mux.HandleFunc(pat.Get("/genre/genres"), getAllGenres)
	mux.HandleFunc(pat.Get("/genre/id/:id"), getByGenereID)
	mux.HandleFunc(pat.Get("/genre/summary"), getGenresSummary)
	mux.HandleFunc(pat.Get("/songs/length/:min/:max"), getByLength)

	http.ListenAndServe("localhost:8080", mux)

}
