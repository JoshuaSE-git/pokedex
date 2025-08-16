package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/JoshuaSE-git/pokedex/internal"
)

const (
	replPrompt  string = "Pokedex > "
	startingURL string = "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
)

func cleanInput(text string) []string {
	lowerCaseInput := strings.ToLower(text)
	splitInput := strings.Fields(lowerCaseInput)

	return splitInput
}

func startRepl() {
	locationUrls := internal.LocationUrls{
		Previous: "",
		Next:     startingURL,
	}
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

		err := command.callback(&locationUrls)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		if command.name == "map" {
			locationUrls, err = internal.GetNextLocationPage(locationUrls.Next)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}

		if command.name == "mapb" && len(locationUrls.Previous) > 0 {
			locationUrls, err = internal.GetNextLocationPage(locationUrls.Previous)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*internal.LocationUrls) error
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
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 locations",
			callback:    commandMapb,
		},
	}
}
