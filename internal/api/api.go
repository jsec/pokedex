package api

type PokeApi struct {
	Location  LocationApi
	Locations LocationsApi
	Pokemon   PokemonApi
}

func NewPokeApi() PokeApi {
	return PokeApi{
		Location:  NewLocationApi(),
		Locations: NewLocationsApi(),
		Pokemon:   NewPokemonApi(),
	}
}
