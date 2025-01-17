package main

import "github.com/rhafaelc/pokedexcli/internal/pokeclient"

func main() {
	client := pokeclient.NewClient()
	pokedex := make(map[string]pokeclient.Pokemon)
	cfg := &config{
		client: client,
		Pokedex: pokedex,
	}
	startRepl(cfg)
}
