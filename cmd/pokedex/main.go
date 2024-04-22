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
	pokeApi := api.NewPokeApi()

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		params := strings.Split(input, " ")
		err := app.ExecCommand(&pokeApi, params)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
