/*
	Ball Clock Main Package
	Brady Clifford
*/

package main

import (
	"fmt"
	"os"
	"flag"
	"BallClock/control"
	)

const MAXBALLS = 127
const MINBALLS = 27

// Gets the paramaters from the command line
// Returns the number of balls to cycle and the number of minutes to run
func parseCommandLine() (int, int) {

	// If no ball count specified, defaults MINBALLS
	ballCount := flag.Int("balls", MINBALLS, fmt.Sprintf("Number of balls to cycle through. Must be between %d and %d.", MINBALLS, MAXBALLS)) // Make sure not an negative value

	// If no minutes to Run, defaults to 0
	minutesToRun := flag.Int("minutes", 0, "Number of minutes to run.") // Make sure not an negative value

	flag.Parse()

	// Ball count must be between the specific range
	if *ballCount < MINBALLS || *ballCount > MAXBALLS {
		fmt.Println("-ball must be between %d and %d.", MINBALLS, MAXBALLS)
		os.Exit(1)
	}

	return *ballCount, *minutesToRun

}

func main() {

	ballCount, minutesToRun := parseCommandLine();

	totalDays := control.Run(uint8(ballCount), uint32(minutesToRun))

	var message string

	if minutesToRun > 0 {
		message = control.GetCurrentStateString()
	} else {
		message = fmt.Sprintf("%d balls cycle after %d days.", ballCount, totalDays)
	}

	fmt.Println(message)
	
}

