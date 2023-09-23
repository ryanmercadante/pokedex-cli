package cli

import (
	"errors"
	"fmt"
)

func explore(cfg *CliConfig, _ ...string) error {
	if cfg.currentLocation == nil {
		return errors.New("travel to a specific location in order to explore")
	}

	locationAreaName := *cfg.currentLocation
	fmt.Printf("Exploring %s...\n", locationAreaName)

	resp, err := cfg.PokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, pokemon := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
