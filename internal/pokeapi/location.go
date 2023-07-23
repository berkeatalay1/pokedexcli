package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) Explore(locationName string) (RespLocationDetails, error) {
	url := baseURL + "/location-area/" + locationName

	cache, isFound := c.cache.Get(url)
	if isFound {
		locationsResp := RespLocationDetails{}
		err := json.Unmarshal(cache, &locationsResp)
		if err != nil {
			return RespLocationDetails{}, err
		}
		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationDetails{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationDetails{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return RespLocationDetails{}, errors.New("Requested Location is not found. Please make sure that location is same as in list")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationDetails{}, err
	}

	c.cache.Add(url, data)

	locationResp := RespLocationDetails{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return RespLocationDetails{}, err
	}

	return locationResp, nil
}
