package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
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
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
