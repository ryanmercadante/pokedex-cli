package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	// check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		locationsAreasResp := LocationAreasResp{}
		err := json.Unmarshal(data, &locationsAreasResp)
		if err != nil {
			return LocationAreasResp{}, err
		}
		return locationsAreasResp, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResp{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResp{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResp{}, fmt.Errorf("bas status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResp{}, err
	}

	locationsAreasResp := LocationAreasResp{}
	err = json.Unmarshal(data, &locationsAreasResp)
	if err != nil {
		return LocationAreasResp{}, err
	}

	c.cache.Add(fullURL, data)

	return locationsAreasResp, nil
}
