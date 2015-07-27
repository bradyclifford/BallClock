/*
A track or rail of a ball clock
*/

package track

import (
	"fmt"
	"encoding/json"
)

type Track struct {
	name string
	// How many balls the track can hold
	capacity int
	// Multiplication value that each ball represents as a minute
	minuteRatio int
	// Slice of balls
	balls []int
}

// Creates a new Track instance
func NewTrack(name string, capacity int, minuteRatio int) Track {

	// Do a one liner here
	if minuteRatio == 0 {
		minuteRatio = 1
	}

	balls := []int{}

	return Track{name, capacity, minuteRatio, balls}

}

// Empty the ball holder and return a reversed list of the flushed balls
func (t *Track) Flush() []int {

	flushedBalls := make([]int, t.capacity)

	for i := range t.balls {
		flushedBalls[(t.capacity - 1) - i] = t.balls[i]
	}

	t.balls = []int{}

	return flushedBalls

}

// Add a ball to the rail.  If the rail is full, it will spill.
// A slice of spilled balls is returned.
func (t *Track) AddBall(ball int) []int {

	if t.IsAtCapacity() {
		return t.Flush()
	}

	t.balls = append(t.balls, ball)
	// Return an empty slice
	return []int{}
}

func (t *Track) IsAtCapacity() bool {
	return len(t.balls) == t.capacity
}

func (t *Track) GetMinutes() int {
	return len(t.balls) * t.minuteRatio
}

func (t *Track) String() string {
	json, _ := json.Marshal(t.balls)
	return fmt.Sprintf("\"%s\":%s", t.name, json)
}