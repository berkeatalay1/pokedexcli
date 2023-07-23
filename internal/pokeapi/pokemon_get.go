package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) Catch(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	cache, isFound := c.cache.Get(url)
	if isFound {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(cache, &pokemonResp)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return Pokemon{}, errors.New("Requested Pokemon is not found. Please make sure that pokemon is same as in list")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, data)

	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemonResp, nil
}
