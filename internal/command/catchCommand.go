package command

import (
	"fmt"
	"math/rand"

	"github.com/jmaeagle99/pokedexcli/internal/pokeapi"
)

func catchCallback(config *CommandConfig, args []string) error {
	client := pokeapi.NewApiClient(&config.Cache)

	pokemon, err := client.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokéball at %s...\n", pokemon.Name)

	chance := int(rand.Float64() * float64(pokemon.BaseExperience))
	if chance < 70 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		config.Caught[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func NewCatchCommand() CliCommand {
	return CliCommand{
		Name:        "catch",
		Description: "Catch a Pokémon",
		Callback:    catchCallback,
	}
}
