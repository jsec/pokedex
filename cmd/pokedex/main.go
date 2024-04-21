package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jsec/pokedex/internal/api"
	"github.com/jsec/pokedex/internal/app"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	pokeApi := api.Create("https://pokeapi.co/api/v2/location-area")

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		err := app.ExecCommand(&pokeApi, input)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
