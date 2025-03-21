package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nonezerone/pokedex-cli/internal/pokeapi"
)

const (
    promptString string = "Pokedex > "
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

type config struct {
    pokeapiClient    pokeapi.Client
    pokedex          pokeapi.Pokedex
    nextLocation     *string
    previousLocation *string
}

func startRepl(cfg *config) {
    scanner := bufio.NewScanner(os.Stdin)
    for {
        fmt.Print(promptString)
        scanner.Scan()

        input := cleanInput((scanner.Text()))
        if len(input) < 1 {
            continue
        }

        commandName, argument := parseArgs(input)

        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback(cfg, argument)
            if err != nil {
                fmt.Println(err)
            }
            continue
        } else {
            fmt.Println("Unknown command")
            continue
        }
    }
}


func cleanInput(text string) []string {
    words := strings.Fields(strings.ToLower(text))
    return words
}

func parseArgs(input []string) (string, string) {
    cmd := input[0]
    arg := ""
    if len(input) > 1 {
        arg = input[1]
    }
    return cmd, arg
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
            callback:    commandMapf,
        },
        "mapb": {
            name:        "mapb",
            description: "Displays the names of 20 previous location areas in the Pokemon world",
            callback:    commandMapb,
        },
        "explore": {
            name:        "explore <location-name>",
            description: "Displays area-specific pokemon",
            callback:    commandExplore,
        },
        "catch": {
            name:        "catch <pokemon>",
            description: "Catches given pokemon (or not)",
            callback:    commandCatch,
        },
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
