/* 
	BallClock Control Package: the state controller
	Keeps track of the cycles performed and the current state of the clock.
	Brady Clifford - July 26, 2015
*/

package control

import (
	"fmt"
	"strings"
	"math"
	"errors"
	"BallClock/track"
	"BallClock/queue"
)

// Static track capacities
const TRACK_MINUTE_NAME = "Min"
const TRACK_MINUTE_CAPACITY = 4
const TRACK_MINUTE_RATION = 1
const TRACK_FIVE_MINUTE_NAME = "FiveMin"
const TRACK_FIVE_MINUTE_CAPACITY = 11
const TRACK_FIVE_MINUTE_RATIO = 5
const TRACK_HOUR_NAME = "Hour"
const TRACK_HOUR_CAPACITY = 11
const TRACK_HOUR_RATIO = 60

// Global state variables
var (

	tracks []track.Track
	clockQueue queue.Queue

	// One tick cycle is a half a day or 12 hours
	clockCycles int

)

// States the clock cycle
// Takes in the capacity of the clock's queue and the minutes to stope the click.
// queueCapacity must be 1 or greater
// minutesToRun can be 0 or greater
// Returns the total days
func Run(queueCapacity int, minutesToRun int) (int, error) {

	if queueCapacity < 1 {
		return 0, errors.New("parameter: capacity must be 1 or greater")
	}

	if minutesToRun < 0 {
		return 0, errors.New("parameter: minutesToRun must be 0 or greater")
	}

	initClock(queueCapacity)

	for  {

		if minutesToRun != 0 && getTotalMinutes() >= minutesToRun {
			break
		}
		
		ball := clockQueue.GetBall()
		digestBall(ball)

		if (clockQueue.IsReset()) {
			break
		}

	}

	return getTotalDays(), nil

}

// Gets the current state of the clock as a string
// TODO: would like to utilize the ToString() override function here
func GetCurrentStateString() string {

	var jsonTracks = make([]string, len(tracks))

	// Iterate through all tracks and get their current state
	for i, track := range tracks {
		jsonTracks[i] = track.String()
	}

	return fmt.Sprintf("{%s,%s}", strings.Join(jsonTracks, ","), clockQueue.String())
}

// Initializes the Clock with the default states
func initClock(queueCapacity int) {

	clockCycles = 0

	clockQueue = queue.NewQueue(queueCapacity)
	tracks = []track.Track{}

	registerTrack(TRACK_MINUTE_NAME, TRACK_MINUTE_CAPACITY, TRACK_MINUTE_RATION)
	registerTrack(TRACK_FIVE_MINUTE_NAME, TRACK_FIVE_MINUTE_CAPACITY, TRACK_FIVE_MINUTE_RATIO)
	registerTrack(TRACK_HOUR_NAME, TRACK_HOUR_CAPACITY, TRACK_HOUR_RATIO)

}

// Takes the passed in ball and determines which track to place it.  
// When a track is full, it is flushed and added back to the queue.
func digestBall(ball int) {

	for index, _ := range tracks {

		// Add ball and if track full, return balls in reverse order. 
		// The orginal ball added does not return in this slice.
		flushedBalls := tracks[index].AddBall(ball)
		
		// No flushed balls returned, meaning there was room on this track.
		if len(flushedBalls) == 0 {
			break
		}

		// If the last track, make sure to add the 
		// ball back into the queue since it has no where else to go
		if (len(tracks) - 1) == index {
			flushedBalls = append(flushedBalls, ball)
			clockCycles++
		}

		clockQueue.AddBalls(flushedBalls) 

	}

}

// Register the track
func registerTrack(name string, capacity int, multiplier int) {
	tracks = append(tracks, track.NewTrack(name, capacity, multiplier))
}

// Get the total number of days that have been processed.
func getTotalDays() int {

	if clockCycles > 0 {
		return int(math.Floor(float64(clockCycles) / 2.0))
	} else { 
		return 0
	}

}

// Gets the total number of minutes that have been processed
// This includes the days past days cycled and the current minutes stored ball clock tracks
func getTotalMinutes() int {
	// 1440 Minutes in a Day
	return int((float64(clockCycles) / 2.0) * 1440.0) + getCurrentCycledMinutes()
}

// Gets only the current minutes stored in the ball clock tracks
func getCurrentCycledMinutes() int {

	var totalCycledMinutes int

	for _, track := range tracks {
		totalCycledMinutes += track.GetMinutes()
	} 

	return totalCycledMinutes

}