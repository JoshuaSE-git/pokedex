package main

import (
	"fmt"
	"os"

	"github.com/JoshuaSE-git/pokedex/internal"
)

func commandExit(locationUrls *internal.LocationUrls) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
