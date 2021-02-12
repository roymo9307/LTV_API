package main

type songs struct {
	ID     int
	Artist string
	Song   string
	Genre  string
	Length int
}

type SongCollection struct {
}

type SongsCollection struct {
	Songs []songs
}

type genres struct {
	//ID          string
	Name        string
	NumSongs    int //int
	Totallength int
}

type GenresCollection struct {
	Genres []genres
}
