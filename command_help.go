package main

import (
	"fmt"

	"github.com/JoshuaSE-git/pokedex/internal"
)

func commandHelp(locationUrls *internal.LocationUrls) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range getCommands() {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}
	fmt.Println()

	return nil
}
