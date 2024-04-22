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

func getCommandList(a *api.PokeApi, params []string) map[string]command {
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
				return commandMap(&a.Locations)
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Fetches the previous 20 locations from the map",
			callback: func() error {
				return commandMapb(&a.Locations)
			},
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific region",
			callback: func() error {
				return commandExplore(a, params[1])
			},
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon",
			callback: func() error {
				return commandCatch(a, params[1])
			},
		},
		"inspect": {
			name:        "inspect",
			description: "Print details of pokemon in the pokedex",
			callback: func() error {
				return commandInspect(a, params[1])
			},
		},
		"pokedex": {
			name:        "pokedex",
			description: "Prints the name of the pokemon in your pokedex",
			callback: func() error {
				return commandPokedex(a)
			},
		},
	}
}

func ExecCommand(a *api.PokeApi, params []string) error {
	cmd, ok := getCommandList(a, params)[params[0]]
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

	for name, details := range getCommandList(nil, nil) {
		fmt.Println(name, "-", details.description)
	}

	return nil
}

func commandMap(l *api.LocationsApi) error {
	l.GetLocationAreas(l.Next)

	for _, location := range l.LocationAreas {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(l *api.LocationsApi) error {
	l.GetLocationAreas(l.Previous)

	for _, location := range l.LocationAreas {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(a *api.PokeApi, name string) error {
	fmt.Println("Exploring", name, "...")
	details := a.Location.GetLocationDetails(name)

	fmt.Println("Found Pokemon:")
	for _, encounter := range details.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}

func commandCatch(a *api.PokeApi, name string) error {
	a.Pokemon.AddToPokedex(name)
	return nil
}

func commandInspect(a *api.PokeApi, name string) error {
	a.Pokemon.Inspect(name)
	return nil
}

func commandPokedex(a *api.PokeApi) error {
	a.Pokemon.Pokedex()
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
