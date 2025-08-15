package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const replPrompt string = "Pokedex > "

func cleanInput(text string) []string {
	lowerCaseInput := strings.ToLower(text)
	splitInput := strings.Fields(lowerCaseInput)

	return splitInput
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(replPrompt)
		scanner.Scan()

		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display usage",
			callback:    commandHelp,
		},
	}
}
