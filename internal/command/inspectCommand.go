package command

import (
	"fmt"
)

func inspectCallback(config *CommandConfig, args []string) error {
	if pokemon, ok := config.Caught[args[0]]; ok {
		fmt.Println("Name: ", pokemon.Name)
		fmt.Println("Height: ", pokemon.Height)
		fmt.Println("Weight: ", pokemon.Weight)
		fmt.Println("Stats:")
		for index := 0; index < len(pokemon.Stats); index++ {
			fmt.Printf("  - %s: %d\n", pokemon.Stats[index].Stat.Name, pokemon.Stats[index].BaseStat)
		}
		fmt.Println("Types:")
		for index := 0; index < len(pokemon.Types); index++ {
			fmt.Printf("  - %s\n", pokemon.Types[index].Type.Name)
		}
	} else {
		fmt.Println("You have not caught that pokemon")
	}
	return nil
}

func NewInspectCommand() CliCommand {
	return CliCommand{
		Name:        "inspect",
		Description: "Inspect a captured Pokémon",
		Callback:    inspectCallback,
	}
}
