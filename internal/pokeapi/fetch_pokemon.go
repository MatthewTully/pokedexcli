package pokeapi

import (
	"encoding/json"
	"fmt"
)

var pokemonPath = "/pokemon"

func FetchPokemon(name string) (PokemonApiResponse, error) {
	url := baseUrl + pokemonPath + "/" + name

	data, err := Get(url)
	if err != nil {
		fmt.Printf("Error occurred Fetching Pokemon: %v\n", err)
		return PokemonApiResponse{}, err
	}
	pokemonJson := PokemonApiResponse{}
	err = json.Unmarshal(data, &pokemonJson)
	if err != nil {
		fmt.Printf("Error Un-marshalling Pokemon JSON: %v\n", err)
		return pokemonJson, err
	}
	return pokemonJson, nil
}
