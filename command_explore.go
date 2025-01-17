package main

import "fmt"

func commandExplore(c *config, area_name *string) error {
	if area_name == nil {
		return fmt.Errorf("No area name given")
	}
	exploreArea, err := c.client.GetExploreArea(*area_name)
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
