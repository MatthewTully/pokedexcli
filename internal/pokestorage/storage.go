package pokestorage

import (
	"fmt"
	"sync"
	"time"

	"github.com/MatthewTully/pokedexcli/internal/pokeapi"
)

type storageItem struct {
	Val      pokeapi.PokemonApiResponse
	CreateAt time.Time
}

type Storage struct {
	pokemon map[string]storageItem
	mux     *sync.Mutex
}

func (s *Storage) Add(key string, val pokeapi.PokemonApiResponse) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if _, exists := s.pokemon[key]; exists {
		return
	}
	s.pokemon[key] = storageItem{
		Val:      val,
		CreateAt: time.Now(),
	}
	fmt.Printf("%v has been added to the Pokédex!\n", key)
}

func (s *Storage) Get(key string) (pokeapi.PokemonApiResponse, bool) {
	s.mux.Lock()
	defer s.mux.Unlock()
	data, exists := s.pokemon[key]
	if exists {
		return data.Val, exists
	}
	return pokeapi.PokemonApiResponse{}, exists
}

func (s *Storage) GetAll() []storageItem {
	s.mux.Lock()
	defer s.mux.Unlock()
	pokeList := []storageItem{}
	for _, v := range s.pokemon {
		pokeList = append(pokeList, v)
	}
	return pokeList

}

func NewStorage() Storage {
	return Storage{
		pokemon: map[string]storageItem{},
		mux:     &sync.Mutex{},
	}
}
