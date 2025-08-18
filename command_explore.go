package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("incorrect usage. expected: explore [location]")
	}

	locationEncountersResp, err := cfg.pokeapiClient.ListEncounters(args[0])
	if err != nil {
		return err
	}

	for _, encounter := range locationEncountersResp.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
