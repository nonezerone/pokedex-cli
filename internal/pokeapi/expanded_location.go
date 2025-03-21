package pokeapi

import (
	"encoding/json"
	"io"
    "fmt"
	"net/http"
    "errors"
)

func (c *Client) ExpandedLocationQuery(location string) (ExpandedAreaResp, error) {
    locationDetails := ExpandedAreaResp{}

	data, exists := c.cache.Get(location)
	if !exists {
        res, err := http.Get("https://pokeapi.co/api/v2/location-area/" + location)
		if err != nil {
			return locationDetails, err
		}

		data, err = io.ReadAll(res.Body)
		res.Body.Close()

		if res.StatusCode == 404 {
			return locationDetails, errors.New(fmt.Sprintf("Location does not exist"))
		}

		if res.StatusCode > 299 {
			return locationDetails, errors.New(fmt.Sprintf("Response failed with status code %d", res.StatusCode))
		}

		c.cache.Add(location, data)
	}

	err := json.Unmarshal(data, &locationDetails)
	if err != nil {
		return locationDetails, errors.New("Error converting JSON response")
	}

	return locationDetails, nil
}
