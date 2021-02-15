package main

import (
	"database/sql"
	"encoding/json"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func allSongs() []byte {
	query := "Select S.id, S.artist, S.song, G.name as genre, s.length FROM Songs s left join Genres G on G.id = S.genre;"
	return querySongs(query)
}

func bySongID(songID string) []byte {
	query := "Select S.id, S.artist, S.song, G.name as genre, s.length FROM Songs s left join Genres G on G.id = S.genre WHERE s.id = \"" + songID + "\";"
	return querySongs(query)
}

func byArtist(artist string) []byte {
	query := "Select S.id, S.artist, S.song, G.name as genre, s.length FROM Songs s left join Genres G on G.id = S.genre WHERE s.artist = \"" + artist + "\";"
	return querySongs(query)
}

func bySong(song string) []byte {
	query := "Select S.id, S.artist, S.song, G.name as genre, s.length FROM Songs s left join Genres G on G.id = S.genre where S.song = \"" + song + "\";"
	return querySongs(query)
}

func byGenre(genre string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where g.name = \"" + genre + "\";"
	return querySongs(query)
}

func allGenres() []byte {
	query := "select g.ID, g.name from Genres AS G;"
	return queryAllGenres(query)
}

func byGenreID(genreID string) []byte {
	query := "select g.ID, g.name from Genres AS G WHERE G.id = \"" + genreID + "\";"
	return queryAllGenres(query)
}

func genresSummary() []byte {
	query := "select g.ID, g.name as genre, count (*) as Num_Song, sum(s.length) as Total_Length from songs S left join genres AS G ON S.genre=G.ID group by genre;"
	return queryGenres(query)
}

func byLength(min, max string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s left join genres g on g.id = s.genre where s.length between " + min + " and " + max + ";"
	return querySongs(query)
}

func querySongs(queryStmt string) []byte {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(queryStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	collection := []Songs{}
	for rows.Next() {
		var id int
		var artist string
		var song string
		var genre string
		var length int
		err = rows.Scan(&id, &artist, &song, &genre, &length)
		if err != nil {
			log.Fatal(err)
		}

		temp := Songs{
			Id:     id,
			Artist: artist,
			Song:   song,
			Genre:  genre,
			Length: length,
		}
		collection = append(collection, temp)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	y := SongsCollection{
		Songs: collection,
	}

	z, _ := json.Marshal(y)

	return z

}

func queryGenres(queryStmt string) []byte {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(queryStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	collection := []Genres{}
	for rows.Next() {
		var ID int
		var Name string
		var num_songs int
		var total_length int
		err = rows.Scan(&ID, &Name, &num_songs, &total_length)
		if err != nil {
			log.Fatal(err)
		}

		temp := Genres{
			ID:           ID,
			Name:         Name,
			Num_Songs:    num_songs,
			Total_length: total_length,
		}
		collection = append(collection, temp)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	y := GenresCollection{
		Genres: collection,
	}

	z, _ := json.Marshal(y)

	return z
}

func queryAllGenres(queryStmt string) []byte {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	stmt, err := db.Prepare(queryStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()

	collection := []Genres{}
	for rows.Next() {
		var ID int
		var Name string

		err = rows.Scan(&ID, &Name)
		if err != nil {
			log.Fatal(err)
		}

		temp := Genres{
			ID:   ID,
			Name: Name,
		}
		collection = append(collection, temp)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	y := GenresCollection{
		Genres: collection,
	}

	z, _ := json.Marshal(y)

	return z
}
