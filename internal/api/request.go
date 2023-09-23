package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func get(client *Client, url string) ([]byte, error) {
	// check cache before making http request
	if data, ok := client.cache.Get(url); ok {
		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("bas status code: %v", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	client.cache.Add(url, data)

	return data, nil
}

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreas, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	locationsAreas := LocationAreas{}
	data, err := get(c, fullURL)
	if err != nil {
		return LocationAreas{}, err
	}

	err = json.Unmarshal(data, &locationsAreas)
	if err != nil {
		return LocationAreas{}, err
	}

	return locationsAreas, nil
}

func (c *Client) GetLocationArea(area string) (LocationArea, error) {
	endpoint := "/location-area/" + area
	fullURL := baseURL + endpoint

	locationArea := LocationArea{}

	data, err := get(c, fullURL)
	if err != nil {
		return LocationArea{}, err
	}

	err = json.Unmarshal(data, &locationArea)
	if err != nil {
		return LocationArea{}, err
	}

	return locationArea, nil
}

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	pokemon := Pokemon{}
	data, err := get(c, fullURL)
	if err != nil {
		return Pokemon{}, err
	}

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
