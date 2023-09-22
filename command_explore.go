package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, param *string) error {
	if param == nil {
		return errors.New("must provide an area name to explore")
	}

	fmt.Printf("Exploring %s...\n", *param)

	resp, err := cfg.pokeapiClient.ListLocationAreaPokemon(param)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
