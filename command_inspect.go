package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args []string) error {

	if len(args) == 0 {
		return errors.New("inspect requires a pokemon name")
	}

	pokemon, ok := cfg.pokedex[args[0]]

	if !ok {
		return errors.New("you have not caught that pokemon")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)

	fmt.Println("Stats:")

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")

	for _, types := range pokemon.Types {
		fmt.Printf("  - %s\n", types.Type.Name)
	}

	return nil
}
