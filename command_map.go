package main

import (
	"errors"
	"fmt"

	"github.com/matthieukhl/go-pokedex/internal/pokecache"
)

func commandMapf(cfg *config, c *pokecache.Cache) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL, c)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *config, c *pokecache.Cache) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL, c)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}

	return nil
}
