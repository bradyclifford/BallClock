/* 
The controller for the ballclock
*/

package control

import (
	"fmt"
	"strings"
	"BallClock/track"
	"BallClock/queue"
	)

// Static track capacities
const TRACK_MINUTE_NAME = "Min"
const TRACK_MINUTE_CAPACITY = 4
const TRACK_FIVE_MINUTE_NAME = "FiveMin"
const TRACK_FIVE_MINUTE_CAPACITY = 11
const TRACK_HOUR_NAME = "Hour"
const TRACK_HOUR_CAPACITY = 11

var tracks []track.Track
var clockQueue queue.Queue

// One click cycle is a half a day or 12 hours
var clockCycles uint32

func Run(queueCapacity uint8, minutesToRun uint32) uint32 {

	initClock(queueCapacity)

	for  {
		
		ball := clockQueue.GetBall()
		contineCycling := digestBall(ball, minutesToRun)

		if (!contineCycling || clockQueue.IsReset()) {
			break
		}

	}

	return getTotalDays()

}

func initClock(queueCapacity uint8) {

	clockCycles = 0

	clockQueue = queue.NewQueue(queueCapacity)
	tracks = []track.Track{}

	registerTrack(TRACK_MINUTE_NAME, TRACK_MINUTE_CAPACITY, 1)
	registerTrack(TRACK_FIVE_MINUTE_NAME, TRACK_FIVE_MINUTE_CAPACITY, 5)
	registerTrack(TRACK_HOUR_NAME, TRACK_HOUR_CAPACITY, 60)

}

func digestBall(ball uint8, minutesToRun uint32) bool {

	continueCycling := true

	for index, track := range tracks {

		if getTotalMinutes() >= minutesToRun {
			continueCycling = false
			break
		}

		flushedBalls := track.AddBall(ball) // returned in reverse order
		
		if len(flushedBalls) == 0 {
			break
		}

		if (len(tracks) - 1) == index {
			flushedBalls = append(flushedBalls, ball)
			clockCycles++
		}

		clockQueue.AddBalls(flushedBalls) 

	}

	return continueCycling

}

func GetCurrentStateString() string {

	var jsonTracks = make([]string, len(tracks))

	for _, track := range tracks {
		jsonTracks = append(jsonTracks, track.String())
	}

	return fmt.Sprintf("{%s,%s}", strings.Join(jsonTracks, ","), clockQueue)
}

func getTotalDays() uint32 {

	if clockCycles > 0 {
		return clockCycles * 2
	} else { 
		return 0
	}

}

func registerTrack(name string, capacity uint8, multiplier uint8) {
	tracks = append(tracks, track.NewTrack(name, capacity, multiplier))
}

func getTotalMinutes() uint32 {
	return (getTotalDays() * 720) + getCurrentCycleMinutes()
}

func getCurrentCycleMinutes() uint32 {

	var totalCycleMinutes uint32

	for _, track := range tracks {
		totalCycleMinutes += track.GetMinutes()
	} 

	return totalCycleMinutes

}