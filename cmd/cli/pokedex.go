package cli

import "fmt"

func pokedex(cfg *CliConfig, _ ...string) error {
	fmt.Println("\nYour Pokedex:")
	for k := range cfg.CaughtPokemon {
		fmt.Printf(" - %s\n", k)
	}
	fmt.Println()
	return nil
}
