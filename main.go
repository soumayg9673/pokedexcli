package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedox > ")
		if scanner.Scan() && scanner.Text() != "" {
			ci := cleanInput(scanner.Text())
			v, ok := cliRegistry()[ci[0]]
			if !ok {
				fmt.Println("Unknown command")
			} else {
				v.callback()
			}
		}
	}
}
