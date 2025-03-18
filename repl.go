package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

const (
    promptString string = "Pokedex > "
)

func startRepl() {
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
            err := command.callback()
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

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
