/*
The ball clock's queue
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

// Remove and return a ball from the beginning of the queue
func (q *Queue) GetBall() int {
	ball := q.balls[0]
	q.balls = q.balls[1:]
	return ball
}

// Put an array of balls back to the end of the queue
// Note: assumes they have already been reversed
func (q *Queue) AddBalls(balls []int) {
	q.balls = append(q.balls, balls...)
}

// Determines if the balls are in their original position in the queue
func (q *Queue) IsReset() bool {
	
	if !q.IsAtCapacity() {
		return false
	}

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
