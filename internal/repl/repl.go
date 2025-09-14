package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jmaeagle99/pokedexcli/internal/command"
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

func RunRepl() {
	commands := command.NewCommandMap()
	config := command.CommandConfig{
		NextUrl:     "",
		PreviousUrl: "",
	}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			tokens := cleanInput(scanner.Text())
			if command, ok := commands[tokens[0]]; ok {
				err := command.Callback(&config)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			} else {
				fmt.Printf("Unknown command '%s'\n", tokens[0])
			}
		}
	}
}
