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

		fmt.Printf("Your command was: %s\n", commandName)
	}
}
