package main

import "github.com/rhafaelc/pokedexcli/internal/pokeclient"

func main() {
	client := pokeclient.NewClient()
	cfg := &config{
		client: client,
	}
	startRepl(cfg)
}
