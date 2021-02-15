# LTV_API

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
##### b)All Songs: http://localhost:8080/songs
##### c)Songs by Id (2-14): http://localhost:8080/songs/id/:id
##### d)Songs by song artist: http://localhost:8080/songs/artist/:artist
##### e)Songs by song name: http://localhost:8080/songs/song/:song
##### f)Songs by song genre: http://localhost:8080/songs/genre/:genre
##### g)All genres: http://localhost:8080/genre/genres
##### h)Genres by Id: http://localhost:8080/genre/genre/id/:id
##### i)Genres summary: http://localhost:8080/genre/genre/summary
##### j)Songs by length between min and max: http://localhost:8080/songs/length/:min/:max
	
#### Examples
##### a)http://localhost:8080/
##### b)http://localhost:8080/songs
##### c)http://localhost:8080/songs/id/7
##### d)http://localhost:8080/songs/artist/Beatles
##### e)http://localhost:8080/songs/song/Amalie
##### f)http://localhost:8080/songs/genre/Pop
##### g)http://localhost:8080/genre/genres
##### h)http://localhost:8080/genre/id/7
##### i)http://localhost:8080/genre/summary
##### j)http://localhost:8080/songs/length/237/450
	
	
