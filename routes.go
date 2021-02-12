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

func getallsongs(w http.ResponseWriter, r *http.Request) {
	z := allsongs()
	w.Write(z)
}

func getbyartist(w http.ResponseWriter, r *http.Request) {
	var artist = pat.Param(r, "artist")
	response := byartist(artist)
	w.Write(response)
}

func getbysong(w http.ResponseWriter, r *http.Request) {
	var song = pat.Param(r, "song")
	response := bysong(song)
	w.Write(response)
}

func getbygenre(w http.ResponseWriter, r *http.Request) {
	var genre = pat.Param(r, "genre")
	response := bygenre(genre)
	w.Write(response)
}

func getgenressummary(w http.ResponseWriter, r *http.Request) {
	response := genressummary()
	w.Write(response)
}

func getbylength(w http.ResponseWriter, r *http.Request) {
	min := pat.Param(r, "min")
	max := pat.Param(r, "max")
	response := bylength(min, max)
	w.Write(response)
}
