package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/soumayg9673/pokedexcli/internal/pokecache"
	"github.com/soumayg9673/pokedexcli/internal/pokedex/pokemon"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := config{
		next:     "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		cache:    *pokecache.NewCache(5 * time.Second),
		pokemons: make(map[string]pokemon.Pokemon),
	}
	for {
		fmt.Print("Pokedox > ")
		if scanner.Scan() && scanner.Text() != "" {
			ci := cleanInput(scanner.Text())
			v, ok := cliRegistry(&c)[ci[0]]
			if !ok {
				fmt.Println("Unknown command")
			} else {
				switch ci[0] {
				case "explore":
					if len(ci) == 2 {
						c.area = ci[1]
					}
					c.catchPokemon = ""
					c.inspectPokemon = ""
				case "catch":
					if len(ci) == 2 {
						c.catchPokemon = ci[1]
					}
					c.area = ""
					c.inspectPokemon = ""
				case "inspect":
					if len(ci) == 2 {
						c.inspectPokemon = ci[1]
					}
					c.area = ""
					c.catchPokemon = ""
				default:
					c.area = ""
					c.catchPokemon = ""
					c.inspectPokemon = ""
				}
				if err := v.callback(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
