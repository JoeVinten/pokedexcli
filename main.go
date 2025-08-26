package main

import (
	"time"

	"github.com/JoeVinten/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 30*time.Second)
	pokedex := make(map[string]pokeapi.Pokemon)
	cfg := &config{
		pokeapiClient: pokeClient,
		pokedex:       pokedex,
	}
	startRepl(cfg)
}
