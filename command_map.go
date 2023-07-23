package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, commandArgs []string) error {

	result, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)

	if err != nil {
		return err
	}
	for _, loc := range result.Results {
		fmt.Println(loc.Name)
	}

	cfg.nextLocationsURL = result.Next
	cfg.prevLocationsURL = result.Previous
	return nil
}

func commandMapb(cfg *config, commandArgs []string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("You are on First Page")
	}
	result, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)

	if err != nil {
		return err
	}

	for _, loc := range result.Results {
		fmt.Println(loc.Name)
	}

	cfg.nextLocationsURL = result.Next
	cfg.prevLocationsURL = result.Previous
	return nil

}
