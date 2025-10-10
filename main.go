package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	c := config{
		next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
	}
	for {
		fmt.Print("Pokedox > ")
		if scanner.Scan() && scanner.Text() != "" {
			ci := cleanInput(scanner.Text())
			v, ok := cliRegistry(&c)[ci[0]]
			if !ok {
				fmt.Println("Unknown command")
			} else {
				v.callback()
			}
		}
	}
}
