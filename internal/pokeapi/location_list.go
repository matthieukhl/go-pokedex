package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/matthieukhl/go-pokedex/internal/pokecache"
)

// List Locations
func (c *Client) ListLocations(pageURL *string, cache *pokecache.Cache) (RespShallowLocations, error) {
	url := baseURL + "/location-area?offset=0&limit=20"

	if pageURL != nil {
		url = *pageURL
	}

	// Check if data exists in cache
	val, exists := cache.Get(url)
	if exists {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	// Add url & data to cache
	cache.Add(url, dat)

	locationResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationResp, nil
}
