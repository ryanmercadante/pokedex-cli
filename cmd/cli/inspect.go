package cli

import (
	"errors"
	"fmt"
)

func inspect(cfg *CliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide a pokemon to inspect")
	}

	pokemonName := args[0]
	pokemon, ok := cfg.CaughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("you haven't caught a %s yet", pokemonName)
	}

	fmt.Printf("\nName: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf(" - %s\n", t.Type.Name)
	}
	fmt.Println()

	return nil
}
