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

func NewTrack(name string, capacity uint8, multiplier uint8) Track {

	// Do a one liner here
	if multiplier == 0 {
		multiplier = 1
	}

	balls := make([]uint8, capacity)

	return Track{name, capacity, multiplier, balls}
}

// Empty the ball holder and return a reversed list of the spilt Balls
func (t *Track) Flush() []uint8 {

	flushedBalls := make([]uint8, t.capacity)

	for i := range t.balls {
		flushedBalls[(t.capacity - 1) - uint8(i)] = t.balls[i]
	}

	return flushedBalls

}

// Add a ball to the rail.  If the rail is full, it will spill.
// A slice of spilled balls is returned.
func (t *Track) AddBall(ball uint8) []uint8 {

	if t.IsAtCapacity() {
		return t.Flush()
	}
	return []uint8{}
}

func (t *Track) IsAtCapacity() bool {
	return uint8(len(t.balls)) == t.capacity
}

func (t *Track) GetMinutes() uint32 {
	return uint32(len(t.balls) * int(t.minuteRatio))
}

func (t *Track) String() string {
	json, _ := json.Marshal(t.balls)
	return t.name + ":" + string(json)
}