package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {

	if len(args) == 0 {
		return errors.New("catch requires a pokemon name")
	}

	pokemonName := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	pokemonDetails, err := cfg.pokeapiClient.GetPokemonDetails(pokemonName)

	if err != nil {
		return err
	}

	catchRoll := rand.Intn(320) + 1

	if catchRoll > pokemonDetails.BaseExperience {
		fmt.Printf("%s was caught\n", pokemonName)
		cfg.pokedex[pokemonName] = pokemonDetails
	} else {
		fmt.Printf("%s escaped\n", pokemonName)
	}

	return nil
}
