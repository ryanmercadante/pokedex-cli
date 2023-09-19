package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ryanmercadante/pokedex-cli/internal/pokemon"
)

func startRepl() {
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	pokeCfg := pokemon.PokeConfig{
		Next:      "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20",
		Previous:  nil,
		Locations: [20]string{},
	}

	for {
		fmt.Print("Pokedex > ")

		// Read next line from input and exit if there's an error
		if ok := scanner.Scan(); !ok {
			break
		}

		// Get text entered by user and clean input
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		command, exists := getCommands(&pokeCfg)[commandName]
		if exists {
			err := command.callback(&pokeCfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *pokemon.PokeConfig) error
}

func getCommands(cfg *pokemon.PokeConfig) map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas",
			callback:    mapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas",
			callback:    mapb,
		},
	}
}
