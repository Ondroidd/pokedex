package main

import (
	"fmt"
	"math/rand/v2"
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

func commandExit(pokemon_data *API_locations, cache *pokecache.Cache, param string, pokedex map[string]API_pokemon) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(pokemon_data *API_locations, cache *pokecache.Cache, param string, pokedex map[string]API_pokemon) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}

	return nil
}

func commandMap(pokemon_data *API_locations, cache *pokecache.Cache, param string, pokedex map[string]API_pokemon) error {
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

func commandMapb(pokemon_data *API_locations, cache *pokecache.Cache, param string, pokedex map[string]API_pokemon) error {
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

func commandExplore(pokemon_data *API_locations, cache *pokecache.Cache, area string, pokedex map[string]API_pokemon) error {
	explore_url := "https://pokeapi.co/api/v2/location-area/" + area

	location, err := get_request[API_explore]("GET", explore_url, cache)
	if err != nil {
		return err
	}

	for _, pokemon := range location.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

func commandCatch(pokemon_data *API_locations, cache *pokecache.Cache, pokemon string, pokedex map[string]API_pokemon) error {
	catch_url := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	res, err := get_request[API_pokemon]("GET", catch_url, cache)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	chance_to_catch := rand.IntN(635)

	if chance_to_catch > (res.BaseExperience * 2) {
		fmt.Printf("%s was caught!\n", pokemon)
		_, ok := pokedex[pokemon]
		if !ok {
			pokedex[pokemon] = res
		}
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemon)

	return nil
}

func commandInspect(pokemon_data *API_locations, cache *pokecache.Cache, pokemon string, pokedex map[string]API_pokemon) error {
	p, ok := pokedex[pokemon]
	if !ok {
		fmt.Printf("%s has not been caught just yet...\n", pokemon)
		return nil
	}

	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
	fmt.Printf("Weight: %d\n", p.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range p.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, pokemon_type := range p.Types {
		fmt.Printf("  - %s\n", pokemon_type.Type.Name)
	}

	return nil
}
