/* 
The controller for the ballclock
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

var tracks []track.Track
var clockQueue queue.Queue

// One click cycle is a half a day or 12 hours
var clockCycles int

func Run(queueCapacity int, minutesToRun int) (int, error) {

	// if queueCapacity == 0, throw an error

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
		digestBall(ball, minutesToRun)

		if (clockQueue.IsReset()) {
			break
		}

	}

	return getTotalDays(), nil

}

func GetCurrentStateString() string {

	var jsonTracks = make([]string, len(tracks))

	for i, track := range tracks {
		jsonTracks[i] = track.String()
	}

	return fmt.Sprintf("{%s,%s}", strings.Join(jsonTracks, ","), clockQueue.String())
}

func initClock(queueCapacity int) {

	clockCycles = 0

	clockQueue = queue.NewQueue(queueCapacity)
	tracks = []track.Track{}

	registerTrack(TRACK_MINUTE_NAME, TRACK_MINUTE_CAPACITY, TRACK_MINUTE_RATION)
	registerTrack(TRACK_FIVE_MINUTE_NAME, TRACK_FIVE_MINUTE_CAPACITY, TRACK_FIVE_MINUTE_RATIO)
	registerTrack(TRACK_HOUR_NAME, TRACK_HOUR_CAPACITY, TRACK_HOUR_RATIO)

}

func digestBall(ball int, minutesToRun int) {

	for index, _ := range tracks {

		// Add ball and if track full, return balls in reverse order. 
		// The orginal ball added does not return in this slice.
		flushedBalls := tracks[index].AddBall(ball)
		
		// No flushed balls returned, meaning there was room on this track.
		if len(flushedBalls) == 0 {
			break
		}

		if (len(tracks) - 1) == index {
			flushedBalls = append(flushedBalls, ball)
			clockCycles++
		}

		clockQueue.AddBalls(flushedBalls) 

	}

}

func registerTrack(name string, capacity int, multiplier int) {
	tracks = append(tracks, track.NewTrack(name, capacity, multiplier))
}

func getTotalDays() int {

	if clockCycles > 0 {
		return int(math.Floor(float64(clockCycles) / 2.0))
	} else { 
		return 0
	}

}

func getTotalMinutes() int {
	// 1440 Minutes in a Day
	return int((float64(clockCycles) / 2.0) * 1440.0) + getCurrentCycledMinutes()
}

func getCurrentCycledMinutes() int {

	var totalCycledMinutes int

	for _, track := range tracks {
		totalCycledMinutes += track.GetMinutes()
	} 

	return totalCycledMinutes

}