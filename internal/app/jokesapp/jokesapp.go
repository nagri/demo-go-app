package jokesapp

import (
	"fmt"
	"net/http"

	"github.com/nagri/demo-go-app/internal/pkg/jokesapi"
)

type Jokeinterface interface {
	GetJoke() (*jokesapi.Joke, error)
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "pong")
}

func LameJokeHandler(w http.ResponseWriter, r *http.Request) {
	j := &jokesapi.Joke{}
	FetchJoke(w, j)
}

func FetchJoke(w http.ResponseWriter, j Jokeinterface) {
	joke, err := j.GetJoke()
	if err != nil {
		fmt.Println(err)
	}
	if joke.Type == "single" {
		fmt.Fprintf(w, "<h1>%s!</h1>", joke.Joke)

	} else if joke.Type == "twopart" {
		fmt.Fprintf(w, "<h1>%s!</h1> <br> <h1>%s</h1>", joke.Setup, joke.Delivery)

	}
}

a