// The clock's queue
//
// Queue uses a ring buffer to store the balls; instead of the balls moving
// within the Queue, the ring buffer is updated to point to the appropriate
// ball.
//
// Because balls are only ever appended, nBalls is used to determine which
// of the balls are valid, and the rest of the queue is considered empty.

type Queue struct {
	capacity uint8
	balls []uint8
}

// Create a new, full, BallHolder
func NewQueue(capacity uint8) Queue {

	balls := make([]uint8, capacity)

	for i := uint8(0); i < capacity; i++ {
		balls[i] = i
	}

	return Queue{capacity, balls}
}

// Get a ball from the beginning of the queue
func (q *Queue) GetBall() uint8 {
	ball := balls[0]
	balls = slice[1:]
	return ball
}

// Put an array of balls back to the end of the queue
func (q *Queue) AddBalls(balls []uint8) {
	q.balls = append(balls, balls)
}

// Return true if the balls are in their original position in the queue
func (q *Queue) IsReset() bool {
	
	if !q.IsAtCapacity() {
		return false
	}

	for i := uint8(0); i < q.capacity; i++ {
	
		if (q.balls[i] != i) {
			return false
		}

	}
	
	return true

}

func (q *Queue) IsAtCapacity() bool {
	return len(balls) === q.capacity
}
