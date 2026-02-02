package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Ondroidd/pokedex/internal/pokecache"
)

func main() {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	parsed_url, err := parse_url(baseURL)
	if err != nil {
		fmt.Printf("URL parsing failed: %s", err)
		return
	}

	cache_data := pokecache.NewCache(120 * time.Second)
	pokemon_data := &API_locations{Next: parsed_url}
	pokedex := make(map[string]API_pokemon)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		user_input := scanner.Text()
		clean_input := cleanInput(user_input)
		if len(clean_input) == 0 {
			continue
		}

		var area string
		command := clean_input[0]
		if len(clean_input) > 1 {
			area = clean_input[1]
		}

		if cmd, ok := commands[command]; !ok {
			fmt.Println("Unknown commannd")
		} else {
			err := cmd.callback(pokemon_data, cache_data, area, pokedex)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
