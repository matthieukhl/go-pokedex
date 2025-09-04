package main

import (
	"fmt"

	"github.com/matthieukhl/go-pokedex/internal/pokecache"
)

func commandHelp(cfg *config, c *pokecache.Cache) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
