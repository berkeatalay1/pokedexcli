package main

import (
	"errors"
)

func commandInspect(cfg *config, commandArgs []string) error {
	if len(commandArgs) == 0 {
		return errors.New("Please write a pokemon name to inspect")
	}

	pokemonName := commandArgs[0]

	pokemon, isFound := cfg.pokedex[pokemonName]
	if !isFound {
		return errors.New("You did not catch " + pokemonName)
	}

	pokemon.Details()
	return nil
}
