package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	cmds := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Give help information",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the name of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the name of the previous 20 location areas in the Pokemon world",
			callback:    commandMapB,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)
	cfg := config{}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cmd, exists := cmds[input]
		if exists {
			err := cmd.callback(&cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	fields := strings.Fields(strings.ToLower(text))
	return fields
}

func commandExit(*config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(*config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("help: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")

	return nil
}

func commandMap(c *config) error {
	var endpoint string

	if len(c.Next) == 0 {
		endpoint = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	} else {
		endpoint = c.Next
	}

	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
		return err
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Printf("Previous URL: %s\nNext URL: %s", c.Previous, c.Next)
	for _, r := range c.Results {
		fmt.Println(r.Name)
	}
	return nil
}

func commandMapB(c *config) error {
	if len(c.Previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	res, err := http.Get(c.Previous)
	if err != nil {
		fmt.Println(err)
		return err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and \nbody: %s\n", res.StatusCode, body)
		return err
	}

	if err != nil {
		fmt.Println(err)
		return err
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Previous URL: %s\nNext URL: %s", c.Previous, c.Next)

	for _, r := range c.Results {
		fmt.Println(r.Name)
	}
	return nil
}

type config struct {
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []areaLocation `json:"results"`
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type areaLocation struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
