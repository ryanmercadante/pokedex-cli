package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemonName string) (PokemonResp, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		pokemonResp := PokemonResp{}
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return PokemonResp{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return PokemonResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonResp{}, fmt.Errorf("bas status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResp{}, err
	}

	pokemonResp := PokemonResp{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return PokemonResp{}, err
	}

	c.cache.Add(fullURL, data)

	return pokemonResp, nil
}
