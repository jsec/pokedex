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

type LocationAreasResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationsApi struct {
	Next          *string
	Previous      *string
	LocationAreas []LocationArea
	cache         cache.Cache
}

func NewLocationsApi() LocationsApi {
	baseUrl := "https://pokeapi.co/api/v2/location-area"

	return LocationsApi{
		Next:  &baseUrl,
		cache: *cache.NewCache(20 * time.Second),
	}
}

func (l *LocationsApi) GetLocationAreas(url *string) {
	if url == nil {
		fmt.Println("No more entries")
		return
	}

	var result LocationAreasResponse

	cached, ok := l.cache.Get(url)
	if ok {
		if err := json.Unmarshal(cached, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
			return
		}

		l.Next = result.Next
		l.Previous = result.Previous
		l.LocationAreas = result.Results

		return
	}

	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err.Error())
	}

	if res.StatusCode > 299 {
		fmt.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}

	l.cache.Add(url, body)

	l.Next = result.Next
	l.Previous = result.Previous
	l.LocationAreas = result.Results
}
