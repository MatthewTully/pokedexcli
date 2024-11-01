package main

import (
	"fmt"

	"github.com/MatthewTully/pokedexcli/internal/pokeapi"
)

func fetchArea(cfg *config) {
	area := cfg.area
	fmt.Printf("Exploring %v...\n", *area)
	areaJson, err := pokeapi.FetchArea(*area)
	if err != nil {
		return
	}
	cfg.currentLocation = area
	if len(areaJson.PokemonEncounters) > 0 {
		fmt.Println("Found Pokemon:")
		for _, pokemons := range areaJson.PokemonEncounters {
			fmt.Printf("- %v\n", pokemons.Pokemon.Name)
		}
		cfg.currentLocationPokemon = &areaJson.PokemonEncounters
		return
	}
	fmt.Println("No Pokemon found!")
}
