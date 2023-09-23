package cli

import (
	"errors"
	"fmt"
)

func travel(cfg *CliConfig, args ...string) error {
	if len(args) != 1 {
		return errors.New("must provide and area name to travel to")
	}

	locationAreaName := args[0]
	resp, err := cfg.PokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	cfg.currentLocation = &resp.Name
	fmt.Println(*cfg.currentLocation)

	return nil
}
