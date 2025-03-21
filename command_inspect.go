package main

import (
    "fmt"
    "errors"
)

func commandInspect(cfg *config, pokemonName string) error {
    pokemon, exists := cfg.pokedex.Entry[pokemonName]
    if !exists {
        return errors.New("you have not caught that pokemon")
    }

    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Printf("Height: %d\n", pokemon.Height)
    fmt.Printf("Weight: %d\n", pokemon.Weight)
    fmt.Println("Stats:")
    for _, stat := range pokemon.Stats {
        fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
    }

    fmt.Println("Types:")
    for _, typ := range pokemon.Types {
        fmt.Printf("  - %s\n", typ.Type.Name)
    }
    return nil
}
