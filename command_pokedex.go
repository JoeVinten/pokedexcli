package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args []string) error {

	if len(cfg.pokedex) == 0 {
		return errors.New("you don't have any pokemon in your pokedex")
	}

	fmt.Println("Your Pokedex:")

	for _, pokemon := range cfg.pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
