package pokeclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetLocationArea(url string) (LocationArea, error) {
	res, err := http.Get(url)
	if err != nil {
		return LocationArea{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return LocationArea{}, fmt.Errorf(
			"Response failed with status code: %d and\nbody: %s\n",
			res.StatusCode,
			body,
		)
	}
	if err != nil {
		return  LocationArea{}, err
	}
	locationArea := LocationArea{}
	if err := json.Unmarshal(body, &locationArea); err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

