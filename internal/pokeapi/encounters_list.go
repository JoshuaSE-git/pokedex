package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListEncounters(location string) (RespLocationEncounters, error) {
	url := baseURL + "/location-area/" + location

	if data, ok := c.cache.Get(url); ok {
		locationEncountersResp := RespLocationEncounters{}

		err := json.Unmarshal(data, &locationEncountersResp)
		if err != nil {
			return RespLocationEncounters{}, err
		}

		return locationEncountersResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationEncounters{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationEncounters{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return RespLocationEncounters{}, errors.New("location not found")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespLocationEncounters{}, err
	}

	locationEncountersResp := RespLocationEncounters{}
	err = json.Unmarshal(data, &locationEncountersResp)
	if err != nil {
		return RespLocationEncounters{}, err
	}

	c.cache.Add(url, data)

	return locationEncountersResp, nil
}
