package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func runRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for true {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			text := scanner.Text()
			tokens := cleanInput(text)
			if len(tokens) > 0 {
				fmt.Println("Your command was:", strings.ToLower(tokens[0]))
			}
		}
	}
}