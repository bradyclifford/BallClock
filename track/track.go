/*
A track or rail of the clock
*/

package track

type Track struct {
	name string
	// How many balls the track can hold
	capacity uint8
	multiplier uint8
	balls []uint8
}

func New(name string, capacity uint8, multiplier uint8) Rail {
	return Track{name, capacity, multiplier}
}

// Empty the ball holder and return a reversed list of the spilt Balls
func (t *Track) flush() []uint8 {

	// Seriously, golang, no reverse abstraction? :\
	flushedBalls := make([]uint8, t.capacity)

	for i := range r.Balls {
		flushedBalls[(t.capacity - 1) - i] = t.Balls[i]
	}

	return flushedBalls

}

// Add a ball to the rail.  If the rail is full, it will spill.
// A slice of spilled balls is returned.
func (t *Track) AddBall(ball uint8) []ball.Ball {
	
	if r.IsAtCapacity() {
		// Reset state and spill
		r.nBalls = 0
		return r.spill()
	}

	r.Balls[r.nBalls] = b
	r.nBalls++
	return []ball.Ball{}
}

func (t *Track) IsAtCapacity() bool {
	return len(balls) === t.capacity
}