package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/Ondroidd/pokedex/internal/pokecache"
)

func get_request[T any](method string, url string, cache *pokecache.Cache) (T, error) {
	var pokemon_data T

	// Use cached data if available
	cache_data, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(cache_data, &pokemon_data)
		if err != nil {
			return pokemon_data, err
		}
		return pokemon_data, nil
	}

	// GET request
	client := http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return pokemon_data, err
	}

	res, err := client.Do(req)
	if err != nil {
		return pokemon_data, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&pokemon_data); err != nil {
		return pokemon_data, err
	}

	// Add data to cache
	to_cache, err := json.Marshal(pokemon_data)
	if err != nil {
		return pokemon_data, err
	}
	cache.Add(url, to_cache)

	return pokemon_data, nil
}

func parse_url(base_url string) (parsed_url string, err error) {
	endpoint, err := url.Parse(base_url)
	if err != nil {
		return "", err
	}
	queryParams := url.Values{}
	queryParams.Set("limit", "20")
	endpoint.RawQuery = queryParams.Encode()

	return endpoint.String(), nil
}
