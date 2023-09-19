package main

import (
	"os"

	"github.com/ryanmercadante/pokedex-cli/internal/pokemon"
)

func commandExit(_ *pokemon.PokeConfig) error {
	os.Exit(0)
	return nil
}
