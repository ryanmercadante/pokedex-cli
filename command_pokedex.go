package main

import "fmt"

func commandPokedex(cfg *config, _ ...string) error {
	fmt.Println("\nYour Pokedex:")
	for k := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", k)
	}
	fmt.Println()
	return nil
}
