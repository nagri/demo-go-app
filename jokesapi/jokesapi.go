package jokesapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Joke struct {
	Error    bool   `json:"error,omitempty"`
	Category string `json:"category,omitempty"`
	Type     string `json:"type,omitempty"`
	Setup    string `json:"setup,omitempty"`
	Joke     string `json:"joke,omitempty"`
	Delivery string `json:"delivery,omitempty"`
	Flags    struct {
		Nsfw      bool `json:"nsfw,omitempty"`
		Religious bool `json:"religious,omitempty"`
		Political bool `json:"political,omitempty"`
		Racist    bool `json:"racist,omitempty"`
		Sexist    bool `json:"sexist,omitempty"`
		Explicit  bool `json:"explicit,omitempty"`
	} `json:"flags,omitempty"`
	Safe bool   `json:"safe,omitempty"`
	ID   int    `json:"id,omitempty"`
	Lang string `json:"lang,omitempty"`
}

// This may return offensive jokes
// const jokeAPI = "https://v2.jokeapi.dev/joke/Any"

const jokeAPI = "https://v2.jokeapi.dev/joke/Programming,Miscellaneous,Pun,Spooky,Christmas?blacklistFlags=religious,racist"

func GetJoke() (*Joke, error) {

	req, _ := http.NewRequest("GET", jokeAPI, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	joke := &Joke{}
	err = json.Unmarshal(body, joke)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return joke, nil

}
