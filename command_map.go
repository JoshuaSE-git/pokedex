package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, args ...string) error {
	locations, err := cfg.pokeapiClient.ListLocations(cfg.nextPageURL)
	if err != nil {
		return err
	}

	cfg.nextPageURL = locations.Next
	cfg.previousPageURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	if cfg.previousPageURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.pokeapiClient.ListLocations(cfg.previousPageURL)
	if err != nil {
		return err
	}

	cfg.nextPageURL = locations.Next
	cfg.previousPageURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
