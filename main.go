package main

import (
	"time"

	"github.com/JoshuaSE-git/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       map[string]pokeapi.Pokemon{},
	}
	startRepl(cfg)
}
