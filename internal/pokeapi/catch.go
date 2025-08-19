package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) Catch(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	if pokemonData, ok := c.cache.Get(name); ok {
		pokemon := Pokemon{}

		err := json.Unmarshal(pokemonData, &pokemon)
		if err != nil {
			return Pokemon{}, nil
		}

		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, nil
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, nil
	}
	defer res.Body.Close()

	pokemonData, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, nil
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(pokemonData, &pokemon)
	if err != nil {
		return Pokemon{}, nil
	}

	c.cache.Add(name, pokemonData)

	return Pokemon{}, nil
}
