package main

import "fmt"

func callbackHelp() {
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Println(cmd.name, " - ", cmd.description)
	}
	fmt.Println()
}
