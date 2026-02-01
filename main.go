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

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		user_input := scanner.Text()
		clean_input := cleanInput(user_input)
		if len(clean_input) == 0 {
			continue
		}

		command := clean_input[0]

		if cmd, ok := commands[command]; !ok {
			fmt.Println("Unknown commannd")
		} else {
			err := cmd.callback(pokemon_data, cache_data)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
