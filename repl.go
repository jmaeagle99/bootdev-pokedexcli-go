package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jmaeagle99/pokedexcli/command"
)

func cleanInput(text string) []string {
	parts := strings.Split(text, " ")
	var result []string
	for _, part := range parts {
		if part != "" {
			result = append(result, part)
		}
	}
	return result
}

func createHelpCallback(commands map[string]command.CliCommand) func() error {
	var helpCallback func() error
	helpCallback = func() error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println()
		for key, value := range commands {
			fmt.Printf("%s: %s\n", key, value.Description)
		}
		return nil
	}
	return helpCallback
}

func createHelpCommand(commands map[string]command.CliCommand) command.CliCommand {
	return command.CliCommand {
		Name: "help",
		Description: "Displays a help message",
		Callback: createHelpCallback(commands),
	}
}

func getCommands() map[string]command.CliCommand {
	commands := make(map[string]command.CliCommand)

	exitCommand := command.NewExitCommand()
	commands[exitCommand.Name] = exitCommand

	helpCommand := createHelpCommand(commands)
	commands[helpCommand.Name] = helpCommand

	return commands
}

func runRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for true {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			tokens := cleanInput(scanner.Text())
			if command, ok := commands[tokens[0]]; ok {
				err := command.Callback()
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			} else {
				fmt.Printf("Unknown command '%s'\n", tokens[0])
			}
		}
	}
}