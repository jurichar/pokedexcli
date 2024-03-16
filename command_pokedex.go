package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Your pokedex:")
	for _, pokemonName := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemonName.Name)
	}
	return nil
}
