package main

import (
	"fmt"

	"github.com/JoshuaSE-git/pokedex/internal"
)

func commandMapb(locationUrls *internal.LocationUrls) error {
	if len(locationUrls.Previous) == 0 {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := internal.GetLocations(locationUrls.Previous)
	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}
