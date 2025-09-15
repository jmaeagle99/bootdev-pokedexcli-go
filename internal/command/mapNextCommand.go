package command

import (
	"fmt"

	"github.com/jmaeagle99/pokedexcli/internal/pokeapi"
)

func newMapNextCallback(config *CommandConfig) error {
	client := pokeapi.NewApiClient(&config.Cache)

	result, err := client.GetLocationAreas(config.NextUrl)
	if err != nil {
		return err
	}

	for _, locationArea := range result.Items {
		fmt.Println(locationArea.Name)
	}

	// Preserve next URL if there is no new next URL so it gets the
	// same page of data on the next invocation.
	if len(result.NextUrl) > 0 {
		config.NextUrl = result.NextUrl
	}
	config.PreviousUrl = result.PreviousUrl

	return nil
}

func NewMapNextCommand() CliCommand {
	return CliCommand{
		Name:        "map",
		Description: "Get the next location areas",
		Callback:    newMapNextCallback,
	}
}
