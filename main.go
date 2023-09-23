package main

import (
	"time"

	"github.com/ryanmercadante/pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.PokemonResp
}

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(5 * time.Hour),
		caughtPokemon: make(map[string]pokeapi.PokemonResp),
	}

	startRepl(cfg)
}
