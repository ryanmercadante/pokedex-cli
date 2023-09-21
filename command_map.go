package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config) error {
	locationAreasResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = locationAreasResp.Next
	cfg.prevLocationAreaURL = locationAreasResp.Previous

	fmt.Println("Location Areas:")
	for _, loc := range locationAreasResp.Results {
		fmt.Printf(" - %s\n", loc.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	locationAreasResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaURL = locationAreasResp.Next
	cfg.prevLocationAreaURL = locationAreasResp.Previous

	for _, loc := range locationAreasResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
