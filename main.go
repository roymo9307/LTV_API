package main

import (
	"net/http"

	"goji.io"
	"goji.io/pat"
)

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/"), homePage)

	mux.HandleFunc(pat.Get("/songs"), getallsongs)

	/*mux.HandleFunc(pat.Get("/songs"), getallsongs)

	mux.HandleFunc(pat.Get("/songs/artist/:artist"), getbyartist)
	mux.HandleFunc(pat.Get("/songs/song/:song"), getbysong)
	mux.HandleFunc(pat.Get("/songs/genre/:genre"), getbygenre)

	mux.HandleFunc(pat.Get("/songs/genre"), getgenressummary)
	mux.HandleFunc(pat.Get("/songs/length/:min/:max"), getbylength)
	*/
	http.ListenAndServe("localhost:8000", mux)

}
