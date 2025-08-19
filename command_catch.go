package main

import (
	"errors"
	"fmt"
)

func commandCatch(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("incorrect usage. expected: catch [pokemon]")
	}

	fmt.Printf("Throwing a Pokeball at %v...", args[0])

	pokemon, err := c.pokeapiClient.Catch(args[0])
	if err != nil {
		return err
	}

	return nil
}
