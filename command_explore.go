package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, locationName string) error {
    if locationName == "" {
        return errors.New("No location to provided to explore!")
    }

    fmt.Println("Exploring " + locationName + "...")
    locationResp, err := cfg.pokeapiClient.ExpandedLocationQuery(locationName)
    if err != nil {
        return err
    }

    fmt.Println("Found Pokemon:")
    for _, entry := range locationResp.PokemonEncounters {
        fmt.Println(" - ", entry.Pokemon.Name)
    }
    return nil
}
