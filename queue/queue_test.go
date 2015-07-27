package queue

import "testing"

func TestNewQueue(t *testing.T) {

	// Arrange
	const EXPECTED_CAPACITY = 30

	// Act
	q := NewQueue(EXPECTED_CAPACITY)

	// Assert

	if q.capacity != EXPECTED_CAPACITY {
		t.Errorf("Unexpected capacity (actual %d, expected %d)",
			q.capacity, EXPECTED_CAPACITY)
	}

	if len(q.balls) != EXPECTED_CAPACITY {
		t.Errorf("Unexpected capacity (actual %d, expected %d)",
			len(q.balls), EXPECTED_CAPACITY)
	}

	// Make sure balls are numbered correctly
	for i, ball := range q.balls {

		i++

		if (ball != i) {
			t.Errorf("Unexpected ball id (actual %d, expected %d)",
				ball, i)
		}

	}

}

func TestShouldGetBallAndRemoveFromQueue(t *testing.T) {

	// Arrange
	const EXPECTED_CAPACITY = 30
	q := NewQueue(EXPECTED_CAPACITY)

	// Act
	ball := q.GetBall()

	// Assert
	if ball != 1 {
		t.Errorf("Should return ball 1 [%d]", ball)
	}

	if len(q.balls) != EXPECTED_CAPACITY - 1 {
		t.Errorf("Unexpected capacity (actual %d, expected %d)",
			len(q.balls), EXPECTED_CAPACITY - 1)
	}

	if q.balls[0] != 2 {
		t.Errorf("Unexpected capacity (actual %d, expected 2)", q.balls[0])
	}
	
}

func TestShouldAddBallsBackIntoQueue(t *testing.T) {

	// Arrange
	const EXPECTED_CAPACITY = 30
	q := NewQueue(EXPECTED_CAPACITY)

	flushedBalls := make([]int, 3)
	flushedBalls[0] = q.GetBall()
	flushedBalls[1] = q.GetBall()
	flushedBalls[2] = q.GetBall()

	// Act
	q.AddBalls(flushedBalls)

	// Assert
	if len(q.balls) != EXPECTED_CAPACITY {
		t.Errorf("Unexpected capacity (actual %d, expected %d)",
			len(q.balls), EXPECTED_CAPACITY)
	}

	if q.balls[29] != 3 {
		t.Errorf("Unexpected capacity (actual %d, expected 3)", q.balls[29])
	}

	if q.balls[28] != 2 {
		t.Errorf("Unexpected capacity (actual %d, expected 2)", q.balls[28])
	}

	if q.balls[27] != 1 {
		t.Errorf("Unexpected capacity (actual %d, expected 1)", q.balls[27])
	}

	if q.balls[0] != 4 {
		t.Errorf("Unexpected capacity (actual %d, expected 4)", q.balls[0])
	}

	if q.balls[1] != 5 {
		t.Errorf("Unexpected capacity (actual %d, expected 5)", q.balls[1])
	}

	if q.balls[2] != 6 {
		t.Errorf("Unexpected capacity (actual %d, expected 6)", q.balls[2])
	}
	
}

func TestShouldNotBeResetWhenQueueIsNotAtCapacity(t *testing.T) {

	// Arrange
	const EXPECTED_CAPACITY = 35
	q := NewQueue(EXPECTED_CAPACITY)

	q.GetBall()

	// Act & Assert
	if q.IsReset() {
		t.Errorf("Should not be reset.")
	}

}

func TestShouldNotBeResetWhenNotInOriginalOrder(t *testing.T) {

	// Arrange
	const EXPECTED_CAPACITY = 35
	q := NewQueue(EXPECTED_CAPACITY)

	q.balls[20] = 10
	q.balls[10] = 20

	// Act & Assert
	if q.IsReset() {
		t.Errorf("Should not be reset.")
	}

}

func TestShouldBeResetWhenInOriginalOrder(t *testing.T) {

	// Arrange
	const EXPECTED_CAPACITY = 35
	q := NewQueue(EXPECTED_CAPACITY)

	// Act & Assert
	if !q.IsReset() {
		t.Errorf("Should be reset [%v]", q)
	}

}