package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/soumayg9673/pokedexcli/internal/pokedex/locationareas"
)

func cliRegistry(c *config) map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedox.",
			callback:    c.onCommandExit,
		},
		"help": {
			name:        "help",
			description: "Show help",
			callback:    c.onCommandHelp,
		},
		"map": {
			name:        "map",
			description: "Location area maps next",
			callback:    c.onCommandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Location area maps previous",
			callback:    c.onCommandMapb,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	next     string
	previous string
}

func (c config) onCommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New(cliRegistry(&c)["exit"].name)
}

func (c config) onCommandHelp() error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
	return errors.New(cliRegistry(&c)["help"].name)
}

func (c *config) onCommandMap() error {
	fullurl := c.next
	if fullurl == "" {
		fullurl = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}
	loc, err := locationareas.GetLocationAreas(fullurl)
	if err != nil {
		return err
	}
	c.next = loc.Next
	c.previous = loc.Previous

	loc.PrintLocationAreaResultsName()

	return errors.New(cliRegistry(c)["map"].name)
}

func (c *config) onCommandMapb() error {
	fullurl := c.previous
	if fullurl == "" {
		fullurl = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	}
	loc, err := locationareas.GetLocationAreas(fullurl)
	if err != nil {
		return err
	}
	c.next = loc.Next
	c.previous = loc.Previous

	loc.PrintLocationAreaResultsName()

	return errors.New(cliRegistry(c)["map"].name)
}
