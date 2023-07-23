package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, commandArgs []string) error {

	pokedex := cfg.pokedex
	if pokedex == nil {
		return errors.New("Error Getting Pokedex")
	}

	fmt.Println("Your Pokedex:")
	for pokemon := range pokedex {
		fmt.Println(" - " + pokemon)
	}

	return nil
}
