package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/berkeatalay1/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient    pokeapi.Client
	pokedex          map[string]pokeapi.Pokemon
	nextLocationsURL *string
	prevLocationsURL *string
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		commandArgs := words[1:]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, commandArgs)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Bring 20 Map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Bring Back 20 Map",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore selected location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch selected location",
			callback:    commandCatch,
		},
	}
}
