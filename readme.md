SUMMARY
=======

Simulation of a ball-clock that outputs the number of
days which elapse before the clock completly cycles through and resets to its orginal state.

USAGE
=================
The program takes input from the command line flag paramaters:

	# Run 30 balls through the ball clock cycle
	BallClock.exe -balls=30
	30 balls cycle after 15 days.
	
	# Run 30 balls through the ball clock cycle but stop at 325 minutes
	BallClock.exe -balls=30 -minutes=325
	{"Min":[],"FiveMin":[22,13,25,3,7],"Hour":[6,12,17,4,15],
	"Main":[11,5,26,18,2,30,19,8,24,10,29,20,16,21,28,1,23,14,27,9]}
	
- balls must be between 27 and 127
- minutes must be a positive integer between 0 to 3477600.

TESTS
=================
To run the unit tests:
	go test ./...
