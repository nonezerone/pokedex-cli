package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/nonezerone/pokeapi"
)

const (
    promptString string = "Pokedex > "
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
    pokeapiClient    pokeapi.Client
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

        commandName := input[0]

        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback(cfg)
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
