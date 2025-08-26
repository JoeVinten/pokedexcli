package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {

	if len(args) == 0 {
		return errors.New("explore requires an area name")
	}

	fmt.Printf("Exploring %s...\n", args[0])

	foundPokemon, err := cfg.pokeapiClient.GetAreaDetails(args[0])

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")

	for _, pokemon := range foundPokemon.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
