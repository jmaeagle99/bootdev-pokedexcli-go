package command

import (
	"fmt"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

func createHelpCallback(commandMap map[string]CliCommand) func() error {
	var helpCallback func() error
	helpCallback = func() error {
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
	return CliCommand {
		Name: "help",
		Description: "Displays a help message",
		Callback: createHelpCallback(commandMap),
	}
}

func NewCommandMap() map[string]CliCommand {
	commandMap := make(map[string]CliCommand)

	commands := []CliCommand {
		NewExitCommand(),
		NewMapCommand(),
		createHelpCommand(commandMap),
	}

	for _, command := range commands {
		commandMap[command.Name] = command
	}

	return commandMap
}