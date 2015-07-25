// The clock's queue

package queue

import "encoding/json"

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
	ball := q.balls[0]
	q.balls = q.balls[1:]
	return ball
}

// Put an array of balls back to the end of the queue
func (q *Queue) AddBalls(balls []uint8) {
	q.balls = append(q.balls, balls...)
}

// Return true if the balls are in their original position in the queue
func (q *Queue) IsReset() bool {
	
	if !q.IsAtCapacity() {
		return false
	}

	for i := uint8(0); i < q.capacity; i++ {
	
		if q.balls[i] != i {
			return false
		}

	}
	
	return true

}

func (q *Queue) IsAtCapacity() bool {
	return uint8(len(q.balls)) == q.capacity
}

func (q *Queue) String() string {
	json, _ := json.Marshal(q.balls)
	return "Main:" + string(json)
}
