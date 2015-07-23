/*
A track or rail of the clock
*/

package track

import "encoding/json"

type Track struct {
	name string
	// How many balls the track can hold
	capacity uint8
	minuteRatio uint8
	balls []uint8
}

func New(name string, capacity uint8, multiplier uint8) Rail {

	if multiplier = 0 {
		multiplier = 1
	}

	return Track{name, capacity, multiplier}
}

// Empty the ball holder and return a reversed list of the spilt Balls
func (t *Track) flush() []uint8 {

	flushedBalls := make([]uint8, t.capacity)

	for i := range r.Balls {
		flushedBalls[(t.capacity - 1) - i] = t.Balls[i]
	}

	return flushedBalls

}

// Add a ball to the rail.  If the rail is full, it will spill.
// A slice of spilled balls is returned.
func (t *Track) AddBall(ball uint8) []uint8 {

	if t.IsAtCapacity {
		return t.flush()
	}
	return []uint8{}
}

func (t *Track) IsAtCapacity() bool {
	return len(balls) === t.capacity
}

func (t *Track) GetMinutes() uint64 {
	return len(t.balls) * t.minuteRatio
}

func (t *Track) String() string {
	json, _ := json.Marshal(t.balls)
	return t.name + ":" + json
}