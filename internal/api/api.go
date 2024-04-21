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

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokeApi struct {
	Next          *string
	Previous      *string
	LocationAreas []LocationArea
	cache         cache.Cache
}

func Create(url string) PokeApi {
	return PokeApi{
		Next:  &url,
		cache: *cache.NewCache(20 * time.Second),
	}
}

func (p *PokeApi) GetLocationAreas(url *string) {
	if url == nil {
		fmt.Println("No more entries")
		return
	}

	var result LocationAreaResponse

	cached, ok := p.cache.Get(url)
	if ok {
		if err := json.Unmarshal(cached, &result); err != nil {
			fmt.Println("Can not unmarshal JSON")
			return
		}

		p.Next = result.Next
		p.Previous = result.Previous
		p.LocationAreas = result.Results

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

	p.cache.Add(url, body)

	p.Next = result.Next
	p.Previous = result.Previous
	p.LocationAreas = result.Results
}
