package main

import (
	"time"

	"github.com/ryanmercadante/pokedex-cli/cmd/cli"
	"github.com/ryanmercadante/pokedex-cli/internal/api"
)

func main() {
	cfg := cli.CliConfig{
		PokeapiClient: api.NewClient(5 * time.Hour),
		CaughtPokemon: make(map[string]api.Pokemon),
		PokeballCount: cli.StarterPokeballCount,
	}

	cli.StartCli(&cfg)
}
