/*
	BallClock Main Package
	Brady Clifford - July 26, 2015

	Summary: 
	Simulation of a ball-clock that outputs the number of
	days which elapse before the clock completly cycles through and resets to its orginal state.

	-balls: must be between 27 and 127
	-minutes: must be a positive integer between 0 to 3477600. 
		Max balls allowed are 127, takes 2415 days to cycle the clock, there are 3477600 minutes in 2415 days.
*/

package main

import (
	"fmt"
	"os"
	"flag"
	"BallClock/control"
	)

const MAX_BALLS = 127
const MIN_BALLS = 27
const MAX_MINUTES_TO_RUN = 3477600

// Gets the paramaters from the command line
// Returns the number of balls to cycle and the number of minutes to run
func parseCommandLine() (int, int) {

	// If no ball count specified, defaults MINBALLS
	ballCount := flag.Int("balls", 0, fmt.Sprintf("Number of balls to cycle through. Must be between %d and %d.", MIN_BALLS, MAX_BALLS)) // Make sure not an negative value

	// If no minutes to Run, defaults to 0
	minutesToRun := flag.Int("minutes", 0, "Number of minutes to run.") // Make sure not an negative value

	flag.Parse()

	// Ball count must be between the specific range
	if *ballCount < MIN_BALLS || *ballCount > MAX_BALLS {
		fmt.Printf("-ball must be between %v and %v. Use attribute -h for help.\n", MIN_BALLS, MAX_BALLS)
		os.Exit(1)
	}

	// Ball count must be between the specific range
	if *minutesToRun < 0 || *minutesToRun > MAX_MINUTES_TO_RUN {
		fmt.Printf("-minutes must be between 0 and %v. Use attribute -h for help.\n", MAX_MINUTES_TO_RUN)
		os.Exit(1)
	}

	return *ballCount, *minutesToRun

}

func main() {

	// Get the paramaters from the command line
	ballCount, minutesToRun := parseCommandLine();

	// Run the clock cycle simulation
	totalDays, err := control.Run(ballCount, minutesToRun)

	var message string

	// Determine which message to display
	if err != nil {
		message = fmt.Sprintf("%v", err)
	} else if minutesToRun > 0 {
		message = control.GetCurrentStateString()
	} else {
		message = fmt.Sprintf("%d balls cycle after %d days.", ballCount, totalDays)
	}

	fmt.Println(message)
	
}

