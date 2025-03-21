package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, _ string) error {
    if len(cfg.pokedex.Entry) < 1 {
        return errors.New("You haven't caught any pokemon yet!")
    }
    for name, _ := range cfg.pokedex.Entry {
        fmt.Printf("  - %s\n", name)
    }
    return nil
}
