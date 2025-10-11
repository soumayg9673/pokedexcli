package main

import (
	"fmt"
	"os"

	"github.com/soumayg9673/pokedexcli/internal/pokecache"
	"github.com/soumayg9673/pokedexcli/internal/pokedex/locationareas"
	"github.com/soumayg9673/pokedexcli/internal/pokedex/pokemon"
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
		"explore": {
			name:        "explore",
			description: "Explore location area",
			callback:    c.onCommandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch Pokemon",
			callback:    c.onCommandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a catched pokemon",
			callback:    c.onCommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "get all my pokemons",
			callback:    c.onCommandPokedex,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

type config struct {
	next           string
	previous       string
	area           string
	catchPokemon   string
	inspectPokemon string
	cache          pokecache.Cache
	pokemons       map[string]pokemon.Pokemon
}

func (c *config) onCommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func (c *config) onCommandHelp() error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
	return nil
}

func (c *config) onCommandMap() error {
	fullurl := c.next

	cacheData, ok := c.cache.Get(fullurl)
	if !ok {
		data, err := locationareas.GetLocationAreas(fullurl)
		if err != nil {
			return err
		}

		c.cache.Add(fullurl, data)

		loc, err := locationareas.GetLocationAreasData(data)
		if err != nil {
			return err
		}

		c.next = loc.Next
		c.previous = loc.Previous

		loc.PrintLocationAreaResultsName()
	} else {
		loc, err := locationareas.GetLocationAreasData(cacheData)
		if err != nil {
			return err
		}
		c.next = loc.Next
		c.previous = loc.Previous

		loc.PrintLocationAreaResultsName()
	}

	return nil
}

func (c *config) onCommandMapb() error {
	fullurl := c.previous

	cacheData, ok := c.cache.Get(fullurl)
	if !ok {
		data, err := locationareas.GetLocationAreas(fullurl)
		if err != nil {
			return err
		}

		c.cache.Add(fullurl, data)

		loc, err := locationareas.GetLocationAreasData(data)
		if err != nil {
			return err
		}

		c.next = loc.Next
		c.previous = loc.Previous
	} else {
		loc, err := locationareas.GetLocationAreasData(cacheData)
		if err != nil {
			return err
		}
		c.next = loc.Next
		c.previous = loc.Previous

		loc.PrintLocationAreaResultsName()
	}

	return nil
}

func (c *config) onCommandExplore() error {
	cacheData, ok := c.cache.Get(c.area)
	if !ok {
		data, err := locationareas.GetPokemonFromLocationArea(c.area)
		if err != nil {
			return err
		}

		c.cache.Add(c.area, data)

		loc, err := locationareas.GetPokemonsFromLocationAreaData(data)
		if err != nil {
			return err
		}

		loc.PrintPokemonsFromLocationAreaResult(c.area)
	} else {
		loc, err := locationareas.GetPokemonsFromLocationAreaData(cacheData)
		if err != nil {
			return err
		}

		loc.PrintPokemonsFromLocationAreaResult(c.area)
	}

	return nil
}

func (c *config) onCommandCatch() error {
	fmt.Printf("Throwing a Pokeball at %s...\n", c.catchPokemon)

	if c.catchPokemon == "" {
		return fmt.Errorf("no pokemon to catch")
	}

	data, err := pokemon.GetPokemon(c.catchPokemon)
	if err != nil {
		return err
	}

	pok, err := pokemon.GetPokemonData(data)
	if err != nil {
		return err
	}

	if _, ok := c.pokemons[pok.Name]; ok {
		return fmt.Errorf("pokemon already caught")
	}

	if catchPokemon := pokemon.CatchPokemon(pok.BaseExperience); catchPokemon {
		fmt.Printf("%s was caught!\n", c.catchPokemon)
		c.pokemons[pok.Name] = pok
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Printf("%s escaped!\n", c.catchPokemon)
	}

	return nil
}

func (c *config) onCommandInspect() error {
	pok, ok := c.pokemons[c.inspectPokemon]
	if !ok {
		return fmt.Errorf("you have not caught that pokemon")
	}

	pok.InspectPokemon()
	return nil
}

func (c *config) onCommandPokedex() error {
	if len(c.pokemons) == 0 {
		return fmt.Errorf("no pokemons caught")
	}

	fmt.Println("Your Pokedex:")
	for p := range c.pokemons {
		fmt.Printf("\t-%s\n", p)
	}
	return nil
}
