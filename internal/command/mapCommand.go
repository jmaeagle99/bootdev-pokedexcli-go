package command

import (
	"fmt"
)

func newMapCallback() error {
	fmt.Println("Listing location areas...")
	return nil
}

func NewMapCommand() CliCommand {
	return CliCommand{
		Name: "map",
		Description: "List next 20 location areas",
		Callback: newMapCallback,
	}
}