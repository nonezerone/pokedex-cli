package main

import (
	"time"

    "github.com/nonezerone/pokedex-cli/internal/pokeapi"
)


func main() {
    pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute*5)

    cfg := &config{
        pokeapiClient: pokeClient,
        pokedex: pokeapi.Pokedex{
            Entry: make(map[string]pokeapi.Pokemon),
        },
    }

    startRepl(cfg)
}
