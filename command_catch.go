package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide a pokemon to catch")
	}

	if cfg.currentLocation == nil {
		return errors.New("there are no pokemon to catch in this location area")
	}

	pokemonName := args[0]
	locationAreaName := *cfg.currentLocation

	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	for _, pokemon := range resp.PokemonEncounters {
		if pokemon.Pokemon.Name == pokemonName {
			fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

			pokemon, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
			if err != nil {
				return err
			}

			const threshold = 50
			randNum := rand.Intn(pokemon.BaseExperience)
			if randNum > threshold {
				return fmt.Errorf("failed to catch %s", pokemonName)
			}

			fmt.Printf("%s was caught!\n", pokemonName)
			cfg.caughtPokemon[pokemonName] = pokemon

			return nil
		}
	}

	return errors.New("can only catch pokemon in your location area")
}
