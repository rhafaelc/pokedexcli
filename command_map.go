package main

import (
	"fmt"

)

func commandMap(c *config) error {
	locationsArea, err := c.client.GetLocationsArea(c.Next)
	if err != nil {
		return err
	}
	c.Next = locationsArea.Next
	c.Previous = locationsArea.Previous
	for _, loc := range locationsArea.Results {
		fmt.Println(loc.Name)	
	}
	return nil
}

func commandMapBack(c *config) error {
	if c.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locationsArea, err := c.client.GetLocationsArea(c.Previous)
	if err != nil {
		return err
	}
	c.Next = locationsArea.Next
	c.Previous = locationsArea.Previous
	for _, loc := range locationsArea.Results {
		fmt.Println(loc.Name)	
	}
	return nil
}
