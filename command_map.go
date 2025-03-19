package main

import (
	"fmt"
    "errors"
)

func commandMapf(cfg *config) error {
	locationsResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation     = locationsResp.Next
	cfg.previousLocation = locationsResp.Previous

	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *config) error {
	if cfg.previousLocation == nil {
        return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocation)
	if err != nil {
		return err
	}

	cfg.nextLocation     = locationResp.Next
	cfg.previousLocation = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
