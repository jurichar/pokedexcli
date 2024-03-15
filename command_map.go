package main

import (
	"fmt"
	"log"

	"github.com/jurichar/pokedexcli/internal/pokeapi"
)

func callbackMap() {
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationArea()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
}
