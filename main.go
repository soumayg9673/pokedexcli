package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/soumayg9673/pokedexcli/internal/pokecache"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := config{
		next:  "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
		cache: *pokecache.NewCache(5 * time.Second),
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
				default:
					c.area = ""
				}
				v.callback()
			}
		}
	}
}
