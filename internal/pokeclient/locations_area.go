package pokeclient

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationsArea(pageUrl *string) (LocationsArea, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	data, exists := c.cache.Get(url)
	if !exists {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationsArea{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationsArea{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationsArea{}, err
		}

		c.cache.Add(url, data)
	}
	locationArea := LocationsArea{}
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationsArea{}, err
	}
	return locationArea, nil
}
