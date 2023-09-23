package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, _ ...string) error {
	if cfg.currentLocation == nil {
		return errors.New("travel to a specific location in order to explore")
	}

	locationAreaName := *cfg.currentLocation
	fmt.Printf("Exploring %s...\n", locationAreaName)

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
