package main

import (
	"fmt"
	"net/http"

	"github.com/nagri/demo-go-app/internal/app/jokesapp"
)

func main() {
	http.HandleFunc("/", jokesapp.LameJokeHandler)
	http.HandleFunc("/ping", jokesapp.Ping)
	fmt.Println(http.ListenAndServe(":8080", nil))
}
