package main

import (
	"fmt"
	"os"
	"flag"
	)

const MAXBALLS = 127
const MINBALLS = 27

func parseCommandLine() (int, int) {

	ballCount := flag.Int("balls", MINBALLS, fmt.Sprintf("Number of balls to cycle through. Must be between %d and %d.", MINBALLS, MAXBALLS))
	minutesToRun := flag.Int("minutes", 0, "Number of minutes to run.")

	flag.Parse()

	if *ballCount < MINBALLS || *ballCount > MAXBALLS {
		fmt.Println("-ball must be between %d and %d.", MINBALLS, MAXBALLS)
		os.Exit(1)
	}

	return *ballCount, *minutesToRun

}

func printMessage(ballCount int, minutesToRun int) {

	var message string

	if minutesToRun > 0 {
		message = fmt.Sprintf("%d minutes to run.", minutesToRun)
	} else {
		message = fmt.Sprintf("%d balls cycle after %d days.", ballCount, 25)
	}

	fmt.Println(message)

}

// Ball Clock
func main() {

	ballCount, minutesToRun := parseCommandLine();

	printMessage(ballCount, minutesToRun)
}