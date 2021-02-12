package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func allsongs() []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre;"
	return querySongs(query)
}

func byartist(artist string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where s.artist = \"" + artist + "\";"
	return querySongs(query)
}

func bysong(song string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where s.song = \"" + song + "\";"
	return querySongs(query)
}

func bygenre(genre string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where g.name = \"" + genre + "\";"
	return querySongs(query)
}

func genressummary() []byte {
	query := "select distinct g.name as genre, count(*) as num_songs, sum(s.length) as total_length from songs s inner join genres g on g.id = s.genre group by genre;"
	return queryGenres(query)
}

func bylength(min, max string) []byte {
	query := "select s.id, s.artist, s.song, g.name as genre, s.length from songs s inner join genres g on g.id = s.genre where s.length between " + min + " and " + max + ";"
	return querySongs(query)
}

func querySongs(queryStmt string) []byte {
	os.Remove("./jrdd.db")

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

	collection := []songs{}
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

		temp := songs{
			ID:     id,
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

	collection := []genres{}
	for rows.Next() {
		//var id int
		var genre string
		var numsongs int
		var totallength int
		err = rows.Scan(&genre, &numsongs, &totallength)
		if err != nil {
			log.Fatal(err)
		}

		temp := genres{
			//ID:           id,
			Name:        genre,
			NumSongs:    numsongs,
			Totallength: totallength,
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
