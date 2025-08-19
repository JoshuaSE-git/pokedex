package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JoshuaSE-git/pokedex/internal/pokeapi"
)

const (
	replPrompt string = "Pokedex > "
)

type config struct {
	nextPageURL     *string
	previousPageURL *string
	location        *string
	pokedex         map[string]pokeapi.Pokemon
	pokeapiClient   pokeapi.Client
}

func cleanInput(text string) []string {
	lowerCaseInput := strings.ToLower(text)
	splitInput := strings.Fields(lowerCaseInput)

	return splitInput
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print(replPrompt)
		scanner.Scan()

		cleanedInput := cleanInput(scanner.Text())
		if len(cleanedInput) == 0 {
			continue
		}

		commandName := cleanedInput[0]
		args := cleanedInput[1:]
		command, ok := getCommands()[commandName]

		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
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
		"map": {
			name:        "map",
			description: "Display the next 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Display location's pokemon encounters",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Throw pokeball at pokemon",
			callback:    commandCatch,
		},
	}
}
