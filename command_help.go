package main

import (
	"fmt"

	"github.com/ryanmercadante/pokedex-cli/internal/pokemon"
)

func commandHelp(cfg *pokemon.PokeConfig) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands(cfg) {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
