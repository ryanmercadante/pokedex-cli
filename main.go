package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	// Create a new scanner to read from standard input
	scanner := bufio.NewScanner(os.Stdin)

	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	for {
		fmt.Print("Pokedex > ")

		// Read next line from input
		if !scanner.Scan() {
			// Exit loop if there's an error
			break
		}

		// Get text entered by user
		text := scanner.Text()

		if command, ok := commands[text]; ok {
			if command.name == "exit" {
				break
			}
			command.callback()
		}
	}

	if scanner.Err() != nil {
		fmt.Println("error:", scanner.Err())
	}
}

func commandHelp() error {
	fmt.Println("\nWelcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("\nhelp: Displays a help message")
	fmt.Println("exit: Exit the Pokedex")
	fmt.Println("")
	return nil
}

func commandExit() error {
	return nil
}
