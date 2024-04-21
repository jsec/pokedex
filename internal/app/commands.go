package app

import (
	"errors"
	"fmt"
	"os"

	"github.com/jsec/pokedex/internal/api"
)

type command struct {
	name        string
	description string
	callback    func() error
}

func getCommandList(p *api.PokeApi) map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback: func() error {
				return commandHelp()
			},
		},
		"exit": {
			name:        "exit",
			description: "Exits the program",
			callback: func() error {
				return commandExit()
			},
		},
		"map": {
			name:        "map",
			description: "Fetches the next 20 locations from the map",
			callback: func() error {
				return commandMap(p)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Fetches the previous 20 locations from the map",
			callback: func() error {
				return commandMapb(p)
			},
		},
	}
}

func ExecCommand(p *api.PokeApi, input string) error {
	cmd, ok := getCommandList(p)[input]
	if !ok {
		return errors.New("Invalid command")
	}

	fmt.Println()

	err := cmd.callback()
	if err != nil {
		return err
	}

	fmt.Println()

	return nil
}

func commandHelp() error {
	fmt.Println("Usage:")

	for name, details := range getCommandList(nil) {
		fmt.Println(name, "-", details.description)
	}

	return nil
}

func commandMap(p *api.PokeApi) error {
	p.GetLocationAreas(p.Next)

	for _, location := range p.LocationAreas {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(p *api.PokeApi) error {
	p.GetLocationAreas(p.Previous)

	for _, location := range p.LocationAreas {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
