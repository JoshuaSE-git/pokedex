package main

import (
	"errors"
	"fmt"
	"math/rand"
)

const (
	minExp = 30
	maxExp = 335
)

func commandCatch(c *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("incorrect usage. expected: catch [pokemon]")
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", args[0])

	pokemon, err := c.pokeapiClient.Catch(args[0])
	if err != nil {
		return err
	}

	if tryCatch(pokemon.BaseExperience) {
		fmt.Printf("%v was caught!\n", pokemon.Name)
		c.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%v escaped!\n", pokemon.Name)
	}

	return nil
}

func tryCatch(baseExp int) bool {
	if baseExp < minExp {
		baseExp = minExp
	}
	if baseExp > maxExp {
		baseExp = maxExp
	}

	normalized := float64(baseExp-minExp) / float64(maxExp-minExp)
	catchRate := int((3.0 + (255.0-3.0)*(1.0-normalized)) + 0.5)

	probability := float64(catchRate) / 255.0

	return rand.Float64() < probability
}
