package pokeclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseUrl + "/pokemon/" + name
	if data, exists := c.cache.Get(url); exists {
		pokemon := Pokemon{}
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return Pokemon{}, nil
		}
		return pokemon, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return Pokemon{}, fmt.Errorf("Not found")
	}


	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(url, data)

	pokemon := Pokemon{}
	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, nil
	}
	return pokemon, nil

}
