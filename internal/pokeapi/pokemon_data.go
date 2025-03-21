package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) GetPokemonData(pokemonName string) (Pokemon, error) {
    pokemonData := Pokemon{}

    data, exists := c.cache.Get(pokemonName)
    if !exists {
        res, err := http.Get(baseURL + "/pokemon/"+pokemonName)
        if err != nil {
            return pokemonData, err
        }

        data, err = io.ReadAll(res.Body)
        res.Body.Close()

        if res.StatusCode == 404 {
            return pokemonData, errors.New("Pokemon does not exist")
        }

        if res.StatusCode > 299 {
            return pokemonData, errors.New("Response failed")
        }

        c.cache.Add(pokemonName, data)
    }

    err := json.Unmarshal(data, &pokemonData)
    if err != nil {
        return pokemonData, errors.New("Error converting JSON response")
    }

    return pokemonData, nil
}
