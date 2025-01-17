package main

import (
	"fmt"
	"math"
	"math/rand"
)

func commandCatch(c *config, pokemon_name *string) error {
	if pokemon_name == nil {
		return fmt.Errorf("No pokemon name given")
	}

	pokemon, err := c.client.GetPokemon(*pokemon_name)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", *pokemon_name)
	catchRate := float64(pokemon.BaseExperience) / float64(400)
	catchRate = math.Floor(catchRate*10) / 10
	catch := rand.Float64() > float64(catchRate)
	if catch {
		fmt.Printf("%s was caught!\n", *pokemon_name)
		c.Pokedex[*pokemon_name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", *pokemon_name)	
	}

	return nil
}
