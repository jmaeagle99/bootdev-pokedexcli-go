package command

import (
	"fmt"

	"github.com/jmaeagle99/pokedexcli/internal/pokeapi"
)

func newMapBackCallback(config *CommandConfig) error {
	result, err := pokeapi.GetLocationAreasResult(config.PreviousUrl)
	if err != nil {
		return err
	}

	for _, locationArea := range result.Items {
		fmt.Println(locationArea.Name)
	}

	config.NextUrl = result.NextUrl
	// Preserve previous URL if there is no new previous URL so it gets the
	// same page of data on the next invocation.
	if len(result.PreviousUrl) > 0 {
		config.PreviousUrl = result.PreviousUrl
	}

	return nil
}

func NewMapBackCommand() CliCommand {
	return CliCommand{
		Name:        "mapb",
		Description: "Get the previous location areas",
		Callback:    newMapBackCallback,
	}
}
