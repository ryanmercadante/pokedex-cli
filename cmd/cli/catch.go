package cli

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	pokeballResetTimer   = 1 * time.Minute
	StarterPokeballCount = 5
)

func catch(cfg *CliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide a pokemon to catch")
	}

	if cfg.currentLocation == nil {
		return errors.New("there are no pokemon to catch in this location area")
	}

	if cfg.PokeballCount == 0 {
		return errors.New("you ran out of pokeballs! After " + pokeballResetTimer.String() + " you will have " + strconv.Itoa(StarterPokeballCount) + " more pokeballs")
	}

	pokemonName := args[0]
	locationAreaName := *cfg.currentLocation

	resp, err := cfg.PokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	for _, pokemon := range resp.PokemonEncounters {
		if pokemon.Pokemon.Name == pokemonName {
			fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

			pokemon, err := cfg.PokeapiClient.GetPokemonInfo(pokemonName)
			if err != nil {
				return err
			}

			const threshold = 50
			randNum := rand.Intn(pokemon.BaseExperience)
			cfg.PokeballCount--

			// if PokeballCount is 0 after decrementing, run a go routine to sleep
			// for `pokeballResetTimer` and then reset PokeballCount to `StarterPokeballCount`
			if cfg.PokeballCount == 0 {
				go func() {
					time.Sleep(pokeballResetTimer)
					cfg.PokeballCount = StarterPokeballCount
				}()
			}

			if randNum > threshold {
				return fmt.Errorf("failed to catch %s", pokemonName)
			}

			fmt.Printf("%s was caught!\n", pokemonName)
			cfg.CaughtPokemon[pokemonName] = pokemon

			return nil
		}
	}

	return errors.New("can only catch pokemon in your location area")
}
