package pokeapi

import (
	"encoding/json"
	"net/http"
)

type LocationArea struct {
	Name string `json:"name"`
}

func GetLocationAreasResult(url string) (PageableResult[LocationArea], error) {
	if len(url) == 0 {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	resp, err := http.Get(url)
	if err != nil {
		return PageableResult[LocationArea]{}, err
	}
	defer resp.Body.Close()

	var result PageableResult[LocationArea]
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&result)
	if err != nil {
		return PageableResult[LocationArea]{}, err
	}
	return result, nil
}
