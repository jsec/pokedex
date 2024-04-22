package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/jsec/pokedex/internal/cache"
)

type LocationDetail struct {
	EncounterMethodRates []EncounterMethodRates `json:"encounter_method_rates"`
	GameIndex            int                    `json:"game_index"`
	ID                   int                    `json:"id"`
	Location             Location               `json:"location"`
	Name                 string                 `json:"name"`
	Names                []Names                `json:"names"`
	PokemonEncounters    []PokemonEncounters    `json:"pokemon_encounters"`
}

type EncounterMethod struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Version struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type VersionDetails struct {
	Rate    int     `json:"rate"`
	Version Version `json:"version"`
}

type EncounterMethodRates struct {
	EncounterMethod EncounterMethod           `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Language struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Names struct {
	Language Language `json:"language"`
	Name     string   `json:"name"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Method struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type EncounterDetails struct {
	Chance          int    `json:"chance"`
	ConditionValues []any  `json:"condition_values"`
	MaxLevel        int    `json:"max_level"`
	Method          Method `json:"method"`
	MinLevel        int    `json:"min_level"`
}

type EncounterVersionDetails struct {
	EncounterDetails []EncounterDetails `json:"encounter_details"`
	MaxChance        int                `json:"max_chance"`
	Version          Version            `json:"version"`
}

type PokemonEncounters struct {
	Pokemon        Pokemon          `json:"pokemon"`
	VersionDetails []VersionDetails `json:"version_details"`
}

type LocationApi struct {
	cache cache.Cache
}

func NewLocationApi() LocationApi {
	return LocationApi{
		cache: *cache.NewCache(20 * time.Second),
	}
}

func (l *LocationApi) GetLocationDetails(name string) LocationDetail {
	baseUrl := "https://pokeapi.co/api/v2/location-area"

	if name == "" {
		fmt.Println("No area name provided")
		return LocationDetail{}
	}

	url := fmt.Sprintf("%v/%v", baseUrl, name)

	var result LocationDetail

	cached, ok := l.cache.Get(&url)
	if ok {
		if err := json.Unmarshal(cached, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
			return LocationDetail{}
		}

		return result
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
		return LocationDetail{}
	}

	l.cache.Add(&url, body)
	return result
}
