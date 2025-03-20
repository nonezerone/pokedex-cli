package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationsAreaRes, error) {
    url := baseURL + "/location-area"
    if pageURL != nil {
        url = *pageURL
    }

    if val, ok := c.cache.Get(url); ok {
        locatonsRes := LocationsAreaRes{}
        err := json.Unmarshal(val, &locatonsRes)
        if err != nil {
            return LocationsAreaRes{}, err
        }

        return locatonsRes, nil
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return LocationsAreaRes{}, err
    }

    res, err := c.httpClient.Do(req)
    if err != nil {
        return LocationsAreaRes{}, err
    }
    defer res.Body.Close()

    data, err := io.ReadAll(res.Body)
    if err != nil {
        return LocationsAreaRes{}, err
    }

    locationsRes := LocationsAreaRes{}
    err = json.Unmarshal(data, &locationsRes)
    if err != nil {
        return LocationsAreaRes{}, err
    }

    c.cache.Add(url, data)
    return locationsRes, err
}
