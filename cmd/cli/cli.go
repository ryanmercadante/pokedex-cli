package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ryanmercadante/pokedex-cli/internal/pokeapi"
)

type CliConfig struct {
	PokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	CaughtPokemon       map[string]pokeapi.PokemonResp
	currentLocation     *string
}

func StartCli(cfg *CliConfig) {
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

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
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("invalid command")
		} else {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
