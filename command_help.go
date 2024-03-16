package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Println(cmd.name, " - ", cmd.description)
	}
	fmt.Println()
	return nil
}
