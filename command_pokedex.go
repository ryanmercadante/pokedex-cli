package main

import "fmt"

func commandPokedex(cfg *config, _ ...string) error {
	fmt.Println("Your Pokedex:")
	for k := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}
