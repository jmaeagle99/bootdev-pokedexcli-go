package command

import (
	"fmt"

	"github.com/jmaeagle99/pokedexcli/internal/pokeapi"
)

func newExploreCallback(config *CommandConfig, args []string) error {
	client := pokeapi.NewApiClient(&config.Cache)

	result, err := client.GetLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", result.Name)

	if len(result.Encounters) > 0 {
		fmt.Println("Found Pokémon:")
		for _, encounter := range result.Encounters {
			fmt.Printf("- %s\n", encounter.Pokemon.Name)
		}
	}

	return nil
}

func NewExploreCommand() CliCommand {
	return CliCommand{
		Name:        "explore",
		Description: "Explore an area for Pokémon",
		Callback:    newExploreCallback,
	}
}
