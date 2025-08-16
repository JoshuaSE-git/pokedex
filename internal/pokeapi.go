package internal

import (
	"encoding/json"
	"net/http"
)

type LocationUrls struct {
	Previous string
	Next     string
}

type Location struct {
	Name string
	URL  string
}

func GetLocations(url string) ([]Location, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response struct {
		Results []Location
	}

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Results, nil
}

func GetNextLocationPage(url string) (LocationUrls, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationUrls{}, err
	}
	defer res.Body.Close()

	var newLocationUrls LocationUrls
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&newLocationUrls)
	if err != nil {
		return LocationUrls{}, err
	}

	return newLocationUrls, nil
}
