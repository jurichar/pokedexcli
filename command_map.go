package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.nextLocationURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous
	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationArea(cfg.prevLocationURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	cfg.nextLocationURL = resp.Next
	cfg.prevLocationURL = resp.Previous
	return nil
}
