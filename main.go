package main

import (
	"time"

	"example.com/nonezerone/pokeapi"
)


func main() {
    pokeClient := pokeapi.NewClient(5 * time.Second)
    cfg := &config{
        pokeapiClient: pokeClient,
    }

    startRepl(cfg)
}
