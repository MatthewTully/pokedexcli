package main

import (
	"fmt"
	"time"

	"github.com/MatthewTully/pokedexcli/internal/pokeapi"
)

func printBasicInformation(data pokeapi.PokemonApiResponse) {
	fmt.Printf("\nName: %v\t# %v\n", data.Name, data.ID)
	fmt.Printf("Height: %v\tWeight: %v\n", data.Height, data.Weight)
}

func printAbilities(data pokeapi.PokemonApiResponse) {
	fmt.Println("\nAbilities:")
	for _, v := range data.Abilities {
		fmt.Printf("  - %v\n", v.Ability.Name)
	}
}

func printMoves(data pokeapi.PokemonApiResponse) {
	fmt.Println("\nLearnable Moves:")
	for _, v := range data.Moves {
		fmt.Printf("  - %v\tLearned by:%v\tLevel required:%v\n",
			v.Move.Name, v.VersionGroupDetails[0].MoveLearnMethod.Name,
			v.VersionGroupDetails[0].LevelLearnedAt)
	}
}

func printStats(data pokeapi.PokemonApiResponse) {
	fmt.Println("\nStats:")
	for _, v := range data.Stats {
		fmt.Printf("  - %v: %v\n", v.Stat.Name, v.BaseStat)
	}
}

func printTypes(data pokeapi.PokemonApiResponse) {
	fmt.Println("\nTypes:")
	for _, v := range data.Types {
		fmt.Printf("  - %v\n", v.Type.Name)
	}
}

func inspectPokemon(cfg *config) {
	pokemon := cfg.pokemon

	data, exists := cfg.pokemonStorage.Get(*pokemon)
	if !exists {
		fmt.Printf("You have not caught a %v\n", *pokemon)
		return
	}
	printBasicInformation(data)
	printTypes(data)
	printStats(data)
	printMoves(data)
	printAbilities(data)

}

func viewPokedex(cfg *config) {
	data := cfg.pokemonStorage.GetAll()

	if len(data) == 0 {
		fmt.Println("Your Pokédex is empty! Explore and catch to populate your dex")
		return
	}
	fmt.Println("Your Pokédex:")
	for _, pokemon := range data {
		fmt.Printf(" - %v\t Captured:%v\n", pokemon.Val.Name, pokemon.CreateAt.Format(time.RFC822))
	}

}
