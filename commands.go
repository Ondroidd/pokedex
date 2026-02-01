package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Ondroidd/pokedex/internal/pokecache"
)

func cleanInput(text string) []string {
	text = strings.ToLower(strings.TrimSpace(text))

	if text == "" {
		return []string{}
	}

	return strings.Fields(text)
}

func commandExit(pokemon_data *API_locations, cache *pokecache.Cache, param string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(pokemon_data *API_locations, cache *pokecache.Cache, param string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMap(pokemon_data *API_locations, cache *pokecache.Cache, param string) error {
	locations, err := get_request[API_locations]("GET", pokemon_data.Next, cache)
	if err != nil {
		return err
	}

	pokemon_data.Next = locations.Next
	pokemon_data.Previous = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(pokemon_data *API_locations, cache *pokecache.Cache, param string) error {
	if pokemon_data.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := get_request[API_locations]("GET", pokemon_data.Previous, cache)
	if err != nil {
		return err
	}

	pokemon_data.Previous = locations.Previous
	pokemon_data.Next = locations.Next

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandExplore(pokemon_data *API_locations, cache *pokecache.Cache, param string) error {
	explore_url := "https://pokeapi.co/api/v2/location-area/" + param

	location, err := get_request[API_explore]("GET", explore_url, cache)
	if err != nil {
		return err
	}

	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}
