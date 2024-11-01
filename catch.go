package main

import (
	"fmt"
	"math/rand"

	"github.com/MatthewTully/pokedexcli/internal/pokeapi"
)

// Capture use rand and base experience
/*
Easy to capture example ~30% 0 > base exp < 100:
 Caterpie - Base experience 39
 Abra - Base experience 62


Mid to capture example ~20% 100 > base exp < 200:
 Noctowl - Base experience 162


Hard to capture example ~10% 200 > base exp < 300
 Arcanine - 213

Extreme to capture example ~5% :
 Mew - 300
 Mew Two - 340


 if threshold is 20 for capture, then
*/

const captureThreshold = 18

func attemptCapture(baseExperience int) bool {
	return rand.Intn(baseExperience) < captureThreshold
}

func pokemonInLocation(pokemon string, area []pokeapi.PokemonEncounters) bool {
	for _, p := range area {
		if p.Pokemon.Name == pokemon {
			return true
		}
	}
	return false
}

func capturePokemon(cfg *config) {
	pokemon := cfg.pokemon
	println()
	if cfg.currentLocationPokemon == nil {
		fmt.Println("No Pokémon to capture. Explore a location to find a Pokémon")
		return
	}

	if !pokemonInLocation(*pokemon, *cfg.currentLocationPokemon) {
		fmt.Printf("You can't see a %v in your current location (%v)\n", *pokemon, *cfg.currentLocation)
		return
	}

	pokemonJson, err := pokeapi.FetchPokemon(*pokemon)
	if err != nil {
		return
	}

	fmt.Printf("Attempting to capture the %v\n", pokemonJson.Name)
	fmt.Println("You throw a poké ball...")
	if attemptCapture(pokemonJson.BaseExperience) {
		fmt.Println("Success!")
		fmt.Printf("You've captured the %v!\n", pokemonJson.Name)

		cfg.pokemonStorage.Add(*pokemon, pokemonJson)
		return
	}
	fmt.Printf("Oh no! The %v broke free!\n", pokemonJson.Name)

}
