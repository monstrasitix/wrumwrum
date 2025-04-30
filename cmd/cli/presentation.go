package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Gray    = "\033[37m"
	White   = "\033[97m"
)

func printColor(color, message string) {
	fmt.Printf("%s%s%s", color, message, Reset)
}

func printHeading() {
	printColor(Red, `
------------------------------------------------------------
  _________.__              .__          __                
 /   _____/|__| _____  __ __|  | _____ _/  |_  ___________ 
 \_____  \ |  |/     \|  |  \  | \__  \\   __\/  _ \_  __ \
 /        \|  |  Y Y  \  |  /  |__/ __ \|  | (  <_> )  | \/
/_______  /|__|__|_|  /____/|____(____  /__|  \____/|__|   
        \/          \/                \/                   
------------------------------------------------------------


`)
}

func getInput() string {
	printColor(Yellow, ":> ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}
