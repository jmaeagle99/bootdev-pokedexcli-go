package command

import "fmt"

func pokedexCallback(config *CommandConfig, args []string) error {
	fmt.Println("Your Pokédex:")
	for name := range config.Caught {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}

func NewPokedexCommand() CliCommand {
	return CliCommand{
		Name:        "pokedex",
		Description: "List caught Pokémon",
		Callback:    pokedexCallback,
	}
}
