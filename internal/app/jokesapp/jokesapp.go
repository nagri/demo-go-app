package jokesapp

import (
	"fmt"
	"net/http"

	"github.com/nagri/demo-go-app/internal/pkg/jokesapi"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", "pong")
}

func LameJokeHandler(w http.ResponseWriter, r *http.Request) {
	j := &jokesapi.Joke{}
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
