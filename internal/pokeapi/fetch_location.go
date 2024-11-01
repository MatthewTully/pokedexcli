package pokeapi

import (
	"encoding/json"
	"fmt"
)

var locationPath = "/location-area"

func FetchMap(pageUrl *string) (LocationGroupApiResponse, error) {
	url := baseUrl + locationPath
	if pageUrl != nil {
		url = *pageUrl
	}

	data, err := Get(url)
	if err != nil {
		fmt.Printf("Error occurred Fetching Locations: %v\n", err)
		return LocationGroupApiResponse{}, err
	}

	areaJson := LocationGroupApiResponse{}
	err = json.Unmarshal(data, &areaJson)
	if err != nil {
		fmt.Printf("Error Un-marshalling Locations JSON: %v\n", err)
		return areaJson, err
	}
	return areaJson, nil
}

func FetchArea(area string) (AreaApiResponse, error) {
	url := baseUrl + locationPath + "/" + area

	data, err := Get(url)
	if err != nil {
		fmt.Printf("Error occurred Fetching Area: %v\n", err)
		return AreaApiResponse{}, err
	}

	areaJson := AreaApiResponse{}
	err = json.Unmarshal(data, &areaJson)
	if err != nil {
		fmt.Printf("Error Un-marshalling Area JSON: %v\n", err)
		return areaJson, err
	}
	return areaJson, nil
}
