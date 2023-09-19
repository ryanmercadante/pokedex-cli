package main

import (
	"fmt"

	"github.com/ryanmercadante/pokedex-cli/internal/pokemon"
)

func mapb(cfg *pokemon.PokeConfig) error {
	err := pokemon.GetPreviousLocationAreas(cfg)
	if err != nil {
		return err
	}
	for i := 0; i < len(cfg.Locations); i++ {
		fmt.Println(cfg.Locations[i])
	}
	return nil
}
