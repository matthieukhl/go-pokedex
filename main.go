package main

import (
	"time"

	"github.com/matthieukhl/go-pokedex/internal/pokeapi"
	"github.com/matthieukhl/go-pokedex/internal/pokecache"
)

const (
	interval = 5
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	c := pokecache.NewCache(interval)

	startRepl(cfg, c)
}
