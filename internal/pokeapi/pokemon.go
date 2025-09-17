package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	Stat     NamedApiResource `json:"stat"`
	Effort   int              `json:"effort"`
	BaseStat int              `json:"base_stat"`
}

type PokemonType struct {
	Order int              `json:"slot"`
	Type  NamedApiResource `json:"type"`
}

func (client *PokeApiClient) GetPokemon(name string) (Pokemon, error) {
	url := fmt.Sprintf("%spokemon/%s", client.baseUrl, name)

	var data []byte
	data, ok := client.cache.Get(url)

	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return Pokemon{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return Pokemon{}, err
		}
	}

	var result Pokemon
	err := json.Unmarshal(data, &result)
	if err != nil {
		return Pokemon{}, err
	}

	client.cache.Add(url, data)

	return result, nil
}
