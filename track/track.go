/* 
	BallClock Track Package: track or rail of a ball clock
	Keeps track of the capacity of the track, stores the balls and flushes them when capacity is met
	Brady Clifford - July 26, 2015
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

	// Don't want to ever divide by zero
	// TODO: maybe should through an error here instead
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

	// Reset to an empty slice
	t.balls = []int{}

	return flushedBalls

}

// Add a ball to the track.  If the rail is at capacity, it will flush.
// Note that the ball that was going to be added, is discarded.
// Returns either an empty slice if the ball is 
// added or a reversed slice of the flushed balls.
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