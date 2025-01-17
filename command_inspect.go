package main

import "fmt"

func commandInspect(c *config, pokemon_name *string) error {
	if pokemon_name == nil {
		return fmt.Errorf("No pokemon name given")
	}
	pokemon, exists := c.Pokedex[*pokemon_name]
	if !exists {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v:%v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, val := range pokemon.Types {
		fmt.Printf("  - %v\n", val.Type.Name)
	}
	return nil
}  
