package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("incorrect usage. expected: explore [location]")
	}
	fmt.Printf("Exploring %v...\n", args[0])

	locationEncountersResp, err := cfg.pokeapiClient.ListEncounters(args[0])
	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon:")
	for _, encounter := range locationEncountersResp.PokemonEncounters {
		fmt.Printf(" - %v\n", encounter.Pokemon.Name)
	}

	return nil
}
