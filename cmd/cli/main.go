package main

import (
	"fmt"
	"strings"
)

func matcher(input string) func(string) bool {
	return func(value string) bool {
		return input == value || strings.Contains(input, value)
	}
}

func cmdHelp() {
	fmt.Println("\thelp\t View commands")
	fmt.Println("\texit\t Exit")
}

func cmdGenerate() {
	fmt.Println("Generating something")
}

func main() {
	printHeading()

	for {
		input := getInput()
		match := matcher(input)

		switch true {
		case match("exit"):
			return

		case match("gen"):
			cmdGenerate()

		case match("help"):
			cmdHelp()
		}
	}
}
