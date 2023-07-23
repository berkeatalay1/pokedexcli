package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, commandArgs []string) error {
	if len(commandArgs) == 0 {
		return errors.New("Please write a location name to explore")
	}
	location := commandArgs[0]
	fmt.Println(fmt.Sprintf(`Exploring %s `, location))
	results, err := cfg.pokeapiClient.Explore(location)

	if err != nil {
		return err
	}

	fmt.Println("Found Pokemon: ")
	for _, encounter := range results.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	fmt.Println()
	return nil
}
