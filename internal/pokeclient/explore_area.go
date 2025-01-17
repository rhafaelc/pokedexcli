package pokeclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetExploreArea(area_name string) (ExploreArea, error) {
	url := baseUrl + "/location-area/" + area_name 
	if data, exists := c.cache.Get(url); exists {
		exploreArea := ExploreArea{}
		if err := json.Unmarshal(data, &exploreArea); err != nil {
			return ExploreArea{}, err
		}
		return exploreArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ExploreArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ExploreArea{}, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusNotFound {
		return ExploreArea{}, fmt.Errorf("Not found")
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreArea{}, err
	}

	c.cache.Add(url, data)

	exploreArea := ExploreArea{}
	if err := json.Unmarshal(data, &exploreArea); err != nil {
		return ExploreArea{}, err
	}
	return exploreArea, nil
}
