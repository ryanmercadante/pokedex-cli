package cli

import (
	"errors"
	"fmt"
)

func mapf(cfg *CliConfig, _ ...string) error {
	locationAreasResp, err := cfg.PokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
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

func mapb(cfg *CliConfig, _ ...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	locationAreasResp, err := cfg.PokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
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
