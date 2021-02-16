# LTV_API
#### This solution uses go code, which consists of making a connection to an SQLite database which can be found at this address at this address https://s3.amazonaws.com/bv-challenge/jrdd.db and must download and place in a folder where it can be accessed by the database file, which makes the connection, consumes the data through structures and then sends it to the main, through the controller, exposing it through JSON code, which is a function performed by an API.
#### This API consists of a list of songs with their respective information and genre. To test the data you can import the database in this link https://extendsclass.com/sqlite-browser.html# and thus review and create the queries


## Prerequisites
#### It should be download Golang 1.15.X or later
#### Download the database in this link https://s3.amazonaws.com/bv-challenge/jrdd.db and place it in the root folder or in a folder that can be accessed by the data file
#### About the packages needed to run the solution are in the github.com and vendor folder, there you can find the following packages:
### go-slite3: 
#### to handle the connection with the local database.
### goji.io: 
#### to manage the connections between the main page and the server.
### golang.org: 
#### to make the connection between the data file and the main page through http.


## This project consists of four main files: 

### Model: 
#### This file defines the structures and structure arrangements which are used in the next file.

### Database: 
#### This file defines the communication between the model and the database, thus taking the data and filling in the structures, since the following file is in charge of making the request for this information, which is why it is waiting to be completed.

### Controller: 
#### This file defines the communication between the main page and the data file, it receives a request from the main page, communicates with the data, and then returns the completed request.

### Main: 
#### This is the main file, the program through a web server accesses it by making a request through a url.


## Run this Project on locally:

#### 1)cd "Your path"
#### 2)git clone https://github.com/roymo9307/LTVTR
#### 3)go build
#### 4)./API
#### 5)Open http://localhost on your browser and type:
##### a)Welcome message: http://localhost:8080/
###### Example http://localhost:8080/
##### b)All Songs: http://localhost:8080/songs
###### Example http://localhost:8080/songs
##### c)Songs by Id (2-14): http://localhost:8080/songs/id/:id
###### Example http://localhost:8080/songs/id/7
##### d)Songs by song artist: http://localhost:8080/songs/artist/:artist
###### Example http://localhost:8080/songs/artist/Beatles
##### e)Songs by song name: http://localhost:8080/songs/song/:song
###### Example http://localhost:8080/songs/song/Amalie
##### f)Songs by song genre: http://localhost:8080/songs/genre/:genre
###### Example http://localhost:8080/songs/genre/Pop
##### g)All genres: http://localhost:8080/genre/genres
###### Example http://localhost:8080/genre/genres
##### h)Genres by Id: http://localhost:8080/genre/genre/id/:id
###### Example http://localhost:8080/genre/id/7
##### i)Genres summary: http://localhost:8080/genre/genre/summary
###### Example http://localhost:8080/genre/summary
##### j)Songs by length between min and max: http://localhost:8080/songs/length/:min/:max
###### Example http://localhost:8080/songs/length/237/450
	

	
	
