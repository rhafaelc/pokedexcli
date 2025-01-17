package main

import "fmt"

func commandPokedex(c *config, _ *string) error {
	fmt.Println("Your Pokedex:")
	for k := range c.Pokedex {
		fmt.Printf(" - %v\n", k)
	}
	return nil
}
