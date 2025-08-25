package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JoeVinten/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsURL *string
	prevLocationsURL *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
	config      *config
}

var commands map[string]cliCommand

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
			config:      nil,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
			config:      nil,
		},
		"map": {
			name:        "map",
			description: "Get all the locations of pokemon",
			callback:    commandMapf,
			config:      &config{},
		},
		"mapb": {
			name:        "map",
			description: "Get the previous 20 locations",
			callback:    commandMapb,
			config:      &config{},
		},
	}

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()

		words := cleanInput(scanner.Text())

		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		if command, exists := commands[commandName]; exists {
			if err := command.callback(cfg); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		} else {
			fmt.Println("Unknown command")
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}
}

func cleanInput(text string) []string {
	lowerCased := strings.ToLower(text)
	return strings.Fields(lowerCased)
}
