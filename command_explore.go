package main

import "fmt"

func commandExplore(c *config, area_name *string) error {
	exploreArea, err := c.client.GetExploreArea(area_name)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", *area_name)
	fmt.Println("Found Pokemon: ")
	for _, pokemon := range exploreArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}
