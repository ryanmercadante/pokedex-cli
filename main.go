package main

import (
	"time"

	"github.com/ryanmercadante/pokedex-cli/cmd/cli"
	"github.com/ryanmercadante/pokedex-cli/internal/pokeapi"
)

func main() {
	cfg := cli.CliConfig{
		PokeapiClient: pokeapi.NewClient(5 * time.Hour),
		CaughtPokemon: make(map[string]pokeapi.PokemonResp),
	}

	cli.StartCli(&cfg)
}
