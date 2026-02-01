package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func get_request(method string, url string) (API_locations, error) {
	client := http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return API_locations{}, err
	}

	res, err := client.Do(req)
	if err != nil {
		return API_locations{}, err
	}

	defer res.Body.Close()

	var pokemon_data API_locations

	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&pokemon_data); err != nil {
		return API_locations{}, err
	}

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
