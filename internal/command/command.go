package command

import (
	"fmt"
	"sort"

	"github.com/jmaeagle99/pokedexcli/internal/pokecache"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*CommandConfig) error
}

type CommandConfig struct {
	Cache       pokecache.Cache
	PreviousUrl string
	NextUrl     string
}

func createHelpCallback(commandMap map[string]CliCommand) func(*CommandConfig) error {
	var helpCallback = func(*CommandConfig) error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println()
		for key, value := range commandMap {
			fmt.Printf("%s: %s\n", key, value.Description)
		}
		return nil
	}
	return helpCallback
}

func createHelpCommand(commandMap map[string]CliCommand) CliCommand {
	return CliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    createHelpCallback(commandMap),
	}
}

func NewCommandMap() map[string]CliCommand {
	commandMap := make(map[string]CliCommand)

	commands := []CliCommand{
		NewExitCommand(),
		NewMapBackCommand(),
		NewMapNextCommand(),
		createHelpCommand(commandMap),
	}

	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Name < commands[j].Name
	})

	for _, command := range commands {
		commandMap[command.Name] = command
	}

	return commandMap
}
