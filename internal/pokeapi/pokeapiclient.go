package pokeapi

import "github.com/jmaeagle99/pokedexcli/internal/pokecache"

type PokeApiClient struct {
	baseUrl string
	cache   *pokecache.Cache
}

func NewApiClient(cache *pokecache.Cache) PokeApiClient {
	return PokeApiClient{
		baseUrl: "https://pokeapi.co/api/v2/",
		cache:   cache,
	}
}
