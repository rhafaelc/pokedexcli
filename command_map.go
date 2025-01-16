package main

import (
	"fmt"

	"github.com/rhafaelc/pokedexcli/pokeclient"
)

func commandMap(c *config) error {
	if c.Next == "" {
		fmt.Println("you're on the last page")
		return nil
	}
	locationArea, err := pokeclient.GetLocationArea(c.Next)
	if err != nil {
		return err
	}
	c.Next = locationArea.Next
	c.Previous = locationArea.Previous
	for _, loc := range locationArea.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}

func commandMapBack(c *config) error {
	if c.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	locationArea, err := pokeclient.GetLocationArea(c.Previous)
	if err != nil {
		return err
	}
	c.Next = locationArea.Next
	c.Previous = locationArea.Previous
	for _, loc := range locationArea.Results {
		fmt.Printf("%s\n", loc.Name)
	}
	return nil
}
