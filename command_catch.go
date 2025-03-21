package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/nonezerone/pokedex-cli/internal/pokeapi"
)

func commandCatch(cfg *config, pokemon string) error {
    if pokemon == "" {
        return errors.New("No pokemon name entered")
    }

    pokemonResp, err := cfg.pokeapiClient.GetPokemonData(pokemon)
    if err != nil {
        return err
    }

    fmt.Println("Throwing a Pokeball at " + pokemon + "...")

    if rand.Intn(pokemonResp.BaseExperience) > pokemonResp.BaseExperience / 2 {
        p := make(map[string]pokeapi.Pokemon)
        p[pokemonResp.Name] = pokemonResp
        cfg.pokedex.Entry = p
        fmt.Printf("%s was caught!\n", pokemonResp.Name)
    } else {
        fmt.Printf("%s escaped!\n", pokemonResp.Name)
    }

    return nil
}
