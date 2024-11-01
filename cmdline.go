package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config)
}

func getCliCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays the help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokédex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Fetch list of locations. (Max of 20 locations at a time. Use 'map' again to fetch the next set).",
			callback:    fetchNextMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Fetch the previous set of location areas. (Can only be used after 'map' has been called).",
			callback:    fetchPrevMap,
		},
		"explore": {
			name:        "explore",
			description: "Explore an area. (Expects kebab case area as input. Use 'map' to view valid areas. E.g: explore canalave-city-area)",
			callback:    fetchArea,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to Catch a Pokémon. Successfully captured Pokémon will be added to your Pokédex! (Pokémon must exist in the location you last explored. E.g: catch tentacool)",
			callback:    capturePokemon,
		},
		"inspect": {
			name:        "inspect",
			description: "View a Pokémon's details. You can only inspect a Pokémon in your dex!",
			callback:    inspectPokemon,
		},
		"pokedex": {
			name:        "pokedex",
			description: "View Pokémon in your dex",
			callback:    viewPokedex,
		},
	}
}

func commandHelp(cfg *config) {
	fmt.Println("\nWelcome to the Pokédex!")
	fmt.Println("Usage:")
	fmt.Println("")
	cmdMap := getCliCommands()
	for _, cmd := range cmdMap {
		fmt.Printf("%v: %v\n", cmd.name, cmd.description)
	}
	fmt.Println("")
}

func commandExit(cfg *config) {
	fmt.Println("Closing the Pokédex... Goodbye!")
	os.Exit(0)
}
