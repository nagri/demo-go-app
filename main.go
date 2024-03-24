package main

import (
	"demogoapp/jokesapi"
	"fmt"
	"net/http"
)

func lameJokeHandler(w http.ResponseWriter, r *http.Request) {
	joke, err := jokesapi.GetJoke()

	if err != nil {
		fmt.Println(err)
	}
	if joke.Type == "single" {
		fmt.Fprintf(w, "<h1>%s!</h1>", joke.Joke)

	} else if joke.Type == "twopart" {
		fmt.Fprintf(w, "<h1>%s!</h1> <br> <h1>%s</h1>", joke.Setup, joke.Delivery)

	}

}

func main() {

	http.HandleFunc("/", lameJokeHandler)
	fmt.Println(http.ListenAndServe(":8080", nil))

}
