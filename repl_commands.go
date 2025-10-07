package main

import (
	"errors"
	"fmt"
	"os"
)

func cliRegistry() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedox.",
			callback:    onCommandExit,
		},
		"help": {
			name:        "help",
			description: "Show help",
			callback:    onCommandHelp,
		},
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func onCommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return errors.New(cliRegistry()["exit"].name)
}

func onCommandHelp() error {
	fmt.Print(`Welcome to the Pokedex!
Usage:

help: Displays a help message
exit: Exit the Pokedex
`)
	return errors.New(cliRegistry()["help"].name)
}
