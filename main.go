package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Pokemon struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			Details string `json:"url"`
		} `json:"ability"`
	} `json:"abilities"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
			Details string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func getPokemon(search string) {
	res, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", search));

	if err != nil {
		fmt.Println(err)
	}

	body, err := io.ReadAll(res.Body)
	
	if err != nil {
		// handle error
	}

	pokemon := &Pokemon{}

	json.Unmarshal(body, &pokemon)

	fmt.Println(pokemon)

}

func main() {
	go getPokemon("pikachu")
	go getPokemon("bulbasaur")
	go getPokemon("charmander")
	go getPokemon("squirtle")
	time.Sleep(time.Second)
}