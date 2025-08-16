package main

import (
	"fmt"

	"github.com/JoshuaSE-git/pokedex/internal"
)

func commandMap(locationUrls *internal.LocationUrls) error {
	locations, err := internal.GetLocations(locationUrls.Next)
	if err != nil {
		return err
	}

	for _, location := range locations {
		fmt.Println(location.Name)
	}

	return nil
}
