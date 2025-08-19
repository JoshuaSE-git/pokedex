package main

import "fmt"

func commandPokedex(c *config, args ...string) error {
	if len(c.pokedex) == 0 {
		fmt.Println("Pokedex is empty! Try catch [pokemon]")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for pokemonName := range c.pokedex {
		fmt.Printf(" - %v\n", pokemonName)
	}

	return nil
}
