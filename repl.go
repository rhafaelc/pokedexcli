package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rhafaelc/pokedexcli/internal/pokeclient"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

type config struct {
	client   pokeclient.Client
	Pokedex  map[string]pokeclient.Pokemon
	Next     *string
	Previous *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)

		if len(words) == 0 {
			continue
		}
		cliCmd := words[0]
		var arg1 *string
		arg1 = nil
		if len(words) > 1 {
			arg1 = &words[1]
		}
		if cmd, ok := getCommands()[cliCmd]; ok {
			err := cmd.callback(cfg, arg1)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 location areas in the pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore <area>",
			description: "Displays pokemon in the area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Catch a pokemon and store the information to pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect pokemon that is registered in pokedex",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Check your pokedex",
			callback:    commandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}
