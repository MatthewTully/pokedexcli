package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/MatthewTully/pokedexcli/internal/pokecache"
)

const baseUrl = "https://pokeapi.co/api/v2"
const cacheInterval = 5 * time.Second

var Cache = pokecache.NewCache(cacheInterval)

func Get(url string) ([]byte, error) {
	data, exists := Cache.Get(url)
	if exists {
		return data, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("response failed with a status: %v", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	Cache.Add(url, body)
	return body, nil
}
