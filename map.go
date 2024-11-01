package main

import (
	"fmt"

	"github.com/MatthewTully/pokedexcli/internal/pokeapi"
)

func printMapResult(results []pokeapi.MapResult) {
	for _, area := range results {
		println(area.Name)
	}
}

func fetchNextMap(cfg *config) {
	areaJson, err := pokeapi.FetchMap(cfg.nextLocationURL)
	if err != nil {
		return
	}
	cfg.nextLocationURL = &areaJson.Next
	cfg.previousLocationURL = areaJson.Previous
	printMapResult(areaJson.Results)

}

func fetchPrevMap(cfg *config) {

	if cfg.previousLocationURL == nil {
		fmt.Println("No Previous value. Has map been called?")
		return
	}
	areaJson, err := pokeapi.FetchMap(cfg.previousLocationURL)
	if err != nil {
		return
	}
	cfg.nextLocationURL = &areaJson.Next
	cfg.previousLocationURL = areaJson.Previous
	printMapResult(areaJson.Results)

}
