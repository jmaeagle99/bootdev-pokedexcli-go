package command

import (
	"fmt"
	"os"
)

func newExitCallback() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func NewExitCommand() CliCommand {
	return CliCommand{
		Name: "exit",
		Description: "Exit the Pokedex",
		Callback: newExitCallback,
	}
}