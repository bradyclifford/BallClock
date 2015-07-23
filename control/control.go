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
var queue queue.Queue
var isInitilized bool

// One click cycle is a half a day or 12 hours
var clockCycles uint64
var minutes uint64
var totalMinutesToRun uint64

func Init(queueCapacity uint8, minutesToRun) {

	totalMinutesToRun = minutesToRun
	clockCycles = 0

	queue = queue.NewQueue(queueCapacity)
	tracks = make([]track.Track)

	registerTrack(TRACK_MINUTE_NAME, TRACK_MINUTE_CAPACITY)
	registerTrack(TRACK_FIVE_MINUTE_NAME, TRACK_FIVE_MINUTE_CAPACITY, 5)
	registerTrack(TRACK_HOUR_NAME, TRACK_HOUR_CAPACITY, 60)

	isInitilized = true

}

func registerTrack(name string, capacity uint8, multiplier uint8) {
	tracks = append(tracks, track.New(name, capacity, multiplier))
}

func Run(queueCapacity uint8, minutesToRun uint64) {

	init(queueCapacity, minutesToRun)

	if isInitilized {

		for  {
			
			ball := queue.GetBall()
			contineCycling := digestBall(ball, minutesToRun)

			if (!contineCycling || queue.IsReset()) {
				break
			}

    	}

	} else {
		// Throw error
	}

}

func digestBall(ball) {

	continueCycling := true

	for index, track := range tracks {

		if haveMinutesRunOut() {
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

		queue.AddBalls(flushedBalls) 

	}

	return continueCycling

}



func GetTotalDays() {

	if initilized && clockCycles > 0 {
		return clockCycles * 2
	} else { 
		return 0
	}

}

func getTotalMinutes() {

	var totalMinutes float64
	
	if initilized {

		totalMinutes += GetTotalDays() * 720 // Total Number of minutes in 12 hours
		totalMinutes += getCurrentMinutes()

	} else {
		return 0
	}

	return 

}

func getCurrentMinutes() (float64) {

	var totalMinutes float64

	for _, track := range tracks {
		totalMinutes += track.GetMinutes()
	} 

	return totalMinutes

}

func haveMinutesRunOut() {
	return getTotalMinutes() == 
}

func ToString() (string) {
	return fmt.Sprintf("{%s,%s}", strings.join(tracks, ","), queue.ToString()
}