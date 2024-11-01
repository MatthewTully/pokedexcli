package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MatthewTully/pokedexcli/internal/pokeapi"
	"github.com/MatthewTully/pokedexcli/internal/pokestorage"
)

type config struct {
	area                   *string
	pokemon                *string
	nextLocationURL        *string
	previousLocationURL    *string
	currentLocation        *string
	currentLocationPokemon *[]pokeapi.PokemonEncounters
	pokemonStorage         pokestorage.Storage
}

func startDex(cfg *config) {
	cfg.pokemonStorage = pokestorage.NewStorage()
	for {
		listenForCmd(cfg)
	}
}

func cleanInput(userInput string) []string {
	lower := strings.ToLower(userInput)
	return strings.Fields((lower))

}

func listenForCmd(cfg *config) {
	cmdMap := getCliCommands()
	scanner := bufio.NewScanner(os.Stdin)
	println()
	fmt.Printf("󰐝 Pokédex > ")
	scanner.Scan()
	userCmd := scanner.Text()
	words := cleanInput(userCmd)
	if len(words) == 0 {
		return
	}
	userCmd = words[0]
	cmd, exists := cmdMap[userCmd]
	if !exists {
		fmt.Printf("\n%s is not a recognised command, please use 'help' to see a list of valid commands.\n\n", userCmd)
		return
	}
	if cmd.name == "explore" {
		if len(words) == 1 {
			fmt.Println("Please specify an area to explore! Use 'map' to see possible locations.")
			return
		}
		cfg.area = &words[1]
	}
	if cmd.name == "catch" || cmd.name == "inspect" {
		if len(words) == 1 {
			fmt.Println("Please specify a Pokémon. Use 'explore' to see the valid Pokémon in a location")
			return
		}
		cfg.pokemon = &words[1]
	}
	cmd.callback(cfg)

}
