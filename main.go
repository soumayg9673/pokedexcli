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
			fmt.Printf("Your command was: %s\n", ci[0])
		}
	}
}
