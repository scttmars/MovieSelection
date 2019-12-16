package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

type Movies struct {

	name string
}


func main() {
	var movieList []Movies
	fptr := flag.String("fpath", "/Users/scottmarsden/Documents/movies.txt", "retrieves text file")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		movieList = append(movieList, Movies{s.Text()})
	}

	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}

	rand.Seed(time.Now().UnixNano())
	var numMovies = len(movieList)
	var movieSelection = movieList[rand.Intn(numMovies)]

	fmt.Println("Tonight's movie is", movieSelection )
}
