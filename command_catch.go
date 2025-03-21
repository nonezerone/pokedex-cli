package main

import (
	"errors"
	"fmt"
	"math/rand"
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
        cfg.pokedex.Entry[pokemonResp.Name] = pokemonResp
        fmt.Printf("%s was caught!\n", pokemonResp.Name)
        fmt.Println("You may now inspect it with the inspect command.")
    } else {
        fmt.Printf("%s escaped!\n", pokemonResp.Name)
    }

    return nil
}
