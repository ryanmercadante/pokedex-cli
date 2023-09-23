package main

import (
	"errors"
	"fmt"
)

func commandTravel(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide and area name to travel to")
	}

	locationAreaName := args[0]
	resp, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	cfg.currentLocation = &resp.Name
	fmt.Println(*cfg.currentLocation)

	return nil
}
