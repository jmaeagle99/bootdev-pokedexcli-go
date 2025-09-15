package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationArea struct {
	Name       string             `json:"name"`
	Encounters []PokemonEncounter `json:"pokemon_encounter"`
}

func (client *PokeApiClient) GetLocationArea(name string) (LocationArea, error) {
	url := fmt.Sprintf("%slocation-area/%s", client.baseUrl, name)

	var data []byte
	data, ok := client.cache.Get(url)

	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return LocationArea{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationArea{}, err
		}
	}

	var result LocationArea
	err := json.Unmarshal(data, &result)
	if err != nil {
		return LocationArea{}, err
	}

	client.cache.Add(url, data)

	return result, nil
}

func (client *PokeApiClient) GetLocationAreas(url string) (PageableResult[NamedApiResource], error) {
	if len(url) == 0 {
		url = fmt.Sprintf("%slocation-area", client.baseUrl)
	}

	var data []byte
	data, ok := client.cache.Get(url)

	if !ok {
		resp, err := http.Get(url)
		if err != nil {
			return PageableResult[NamedApiResource]{}, err
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return PageableResult[NamedApiResource]{}, err
		}
	}

	var result PageableResult[NamedApiResource]
	err := json.Unmarshal(data, &result)
	if err != nil {
		return PageableResult[NamedApiResource]{}, err
	}

	client.cache.Add(url, data)

	return result, nil
}
