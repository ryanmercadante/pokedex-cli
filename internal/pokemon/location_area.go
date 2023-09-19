package pokemon

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type PokeConfig struct {
	Next      string
	Previous  *string
	Locations [20]string
}

func GetNextLocationAreas(cfg *PokeConfig) error {
	res, err := http.Get(cfg.Next)
	if err != nil {
		return errors.New("error trying to get next pokemon location areas")
	}
	err = processResponse(res, cfg)
	if err != nil {
		return err
	}
	return nil
}

func GetPreviousLocationAreas(cfg *PokeConfig) error {
	if cfg.Previous == nil {
		return errors.New("there are no previous pokemon location areas")
	}
	res, err := http.Get(*cfg.Previous)
	if err != nil {
		return errors.New("error trying to get next pokemon location areas")
	}
	err = processResponse(res, cfg)
	if err != nil {
		return err
	}
	return nil
}

func processResponse(response *http.Response, cfg *PokeConfig) error {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.New("error trying to read body")
	}
	data := LocationAreaResponse{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return errors.New("error trying to unmarshall json")
	}
	cfg.Next = data.Next
	cfg.Previous = data.Previous
	for i := 0; i < len(data.Results); i++ {
		cfg.Locations[i] = data.Results[i].Name
	}
	return nil
}
