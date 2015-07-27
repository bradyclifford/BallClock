/*
The ball clock's queue
*/

/* 
	BallClock Queue Package: keeps track of queue balls and determines when the queue is reset
	Brady Clifford - July 26, 2015
*/

package queue

import "encoding/json"

type Queue struct {
	capacity int
	balls []int
}

// Create a new, full queue
func NewQueue(capacity int) Queue {

	balls := make([]int, capacity)

	// Iterate through and add number which identifies a ball
	for i := 0; i < capacity; i++ {
		balls[i] = i + 1
	}

	return Queue{capacity, balls}
}

// Return a ball from the beginning of the queue and remove it from the slice
func (q *Queue) GetBall() int {
	ball := q.balls[0]
	q.balls = q.balls[1:]
	return ball
}

// Add a slice of balls to the end of the queue
// Note: assumes they have already been reversed
func (q *Queue) AddBalls(balls []int) {
	q.balls = append(q.balls, balls...)
}

// Determines if the balls are in their original position in the queue
func (q *Queue) IsReset() bool {
	
	// Can't be reset if not at its orginal capacity
	if !q.IsAtCapacity() {
		return false
	}

	// Compare each ball in queue to see if it is in order
	for i := 0; i < q.capacity; i++ {
		if q.balls[i] != (i + 1) {
			return false
		}
	}
	
	return true

}

func (q *Queue) IsAtCapacity() bool {
	return len(q.balls) == q.capacity
}

func (q *Queue) String() string {
	json, _ := json.Marshal(q.balls)
	return "\"Main\":" + string(json)
}
