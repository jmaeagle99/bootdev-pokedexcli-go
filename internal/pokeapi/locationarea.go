package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jmaeagle99/pokedexcli/internal/pokecache"
)

type LocationArea struct {
	Name string `json:"name"`
}

func GetLocationAreasResult(url string, cache *pokecache.Cache) (PageableResult[LocationArea], error) {
	if len(url) == 0 {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	var data []byte
	data, ok := cache.Get(url)

	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return PageableResult[LocationArea]{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return PageableResult[LocationArea]{}, err
		}
	}

	var result PageableResult[LocationArea]
	err := json.Unmarshal(data, &result)
	if err != nil {
		return PageableResult[LocationArea]{}, err
	}

	cache.Add(url, data)

	return result, nil
}
