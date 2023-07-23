package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, commandArgs []string) error {
	if len(commandArgs) == 0 {
		return errors.New("Please write a pokemon name to catch")
	}
	pokemonName := commandArgs[0]
	fmt.Println(fmt.Sprintf(`Throwing a Pokeball at %s...`, pokemonName))
	pokemon, err := cfg.pokeapiClient.Catch(pokemonName)

	if err != nil {
		return err
	}
	requiredPercent := 0.5

	if pokemon.BaseExperience < 100 {
		requiredPercent = 0.8
	} else if pokemon.BaseExperience < 200 {
		requiredPercent = .5
	} else {
		requiredPercent = .3
	}
	if rand.Float64() < requiredPercent {
		fmt.Println(fmt.Sprintf(`%s was caught!!!`, pokemonName))
		cfg.pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Println(fmt.Sprintf(`%s excaped!!!`, pokemonName))
	}

	fmt.Println()
	return nil
}
