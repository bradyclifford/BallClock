package track

import "testing"

func TestNewTrack(t *testing.T) {

	// Arrange

	const EXPECTED_NAME = "Test"
	const EXPECTED_CAPACITY = 11
	const EXPECTED_MINUTE_RATIO = 2

	// Act

	tr := NewTrack(EXPECTED_NAME, EXPECTED_CAPACITY, EXPECTED_MINUTE_RATIO)

	// Assert

	if tr.name != EXPECTED_NAME {
		t.Errorf("Unexpected name (actual %d, expected %d)",
			tr.capacity, EXPECTED_NAME)
	}

	if tr.capacity != EXPECTED_CAPACITY {
		t.Errorf("Unexpected capacity (actual %d, expected %d)",
			tr.capacity, EXPECTED_CAPACITY)
	}

	if tr.minuteRatio != EXPECTED_MINUTE_RATIO {
		t.Errorf("Unexpected minute ratio (actual %d, expected %d)",
			tr.minuteRatio, EXPECTED_MINUTE_RATIO)
	}
	
}

func TestShouldAddBallFirst(t *testing.T) {

	const BALL_1 = 5 

	// Arrange
	tr := NewTrack("Test", 4, 0)

	// Act
	flushedBalls := tr.AddBall(BALL_1)
	actual_ball := tr.balls[0]

	// Assert
	if actual_ball != BALL_1 {
		t.Errorf("Unexpected ball value(actual %d, expected %d)",
			actual_ball, BALL_1)
	}

	if len(flushedBalls) != 0 {
		t.Errorf("No balls should be flushed [%+v]", flushedBalls)
	}

}

func TestShouldAddBallSecond(t *testing.T) {

	const BALL_1 = 5 
	const BALL_2 = 6

	// Arrange
	tr := NewTrack("Test", 4, 0)

	// Act
	flushedBalls := tr.AddBall(BALL_1)
	if len(flushedBalls) != 0 {
		t.Errorf("No balls should be flushed [%+v]", flushedBalls)
	}

	flushedBalls = tr.AddBall(BALL_2)
	if len(flushedBalls) != 0 {
		t.Errorf("No balls should be flushed [%+v]", flushedBalls)
	}

	actual_ball1 := tr.balls[0]
	actual_ball2 := tr.balls[1]

	// Assert
	if actual_ball1 != BALL_1 {
		t.Errorf("Unexpected ball1 value(actual %d, expected %d)",
			actual_ball1, BALL_1)
	}

	if actual_ball2 != BALL_2 {
		t.Errorf("Unexpected ball2 value(actual %d, expected %d)",
			actual_ball2, BALL_2)
	}

}

func TestShouldAddBallAndFlush(t *testing.T) {

	const BALL_1 = 5 
	const BALL_2 = 6
	const BALL_3 = 7

	// Arrange
	tr := NewTrack("Test", 2, 0)

	// Act
	tr.AddBall(BALL_1)
	tr.AddBall(BALL_2)
	flushedBalls := tr.AddBall(BALL_3)

	// Assert
	if tr.IsAtCapacity() {
		t.Errorf("Balls should of been flushed [%+v]", tr)
	}

	if len(flushedBalls) != 2 {
		t.Errorf("Flushed balls should of been returned [%+v]", flushedBalls)
	}

	// Should return balls in reverse order
	if flushedBalls[0] != BALL_2 {
		t.Errorf("Unexpected ball2 value(actual %d, expected %d)",
			flushedBalls[0], BALL_2)
	}

	if flushedBalls[1] != BALL_1 {
		t.Errorf("Unexpected ball2 value(actual %d, expected %d)",
			flushedBalls[1], BALL_1)
	}
	
}

func TestShouldGetMinutes(t *testing.T) {

	const BALL_1 = 5
	const BALL_2 = 6

	// Arrange
	tr := NewTrack("Test", 11, 5) // 5 minute track representation

	// Act
	tr.AddBall(BALL_1)
	tr.AddBall(BALL_2)

	// Assert
	if tr.GetMinutes() != 10 {
		t.Errorf("Should calculate 10 minutes [%d]", tr.GetMinutes())
	}

}

func TestShouldConvertStateToString(t *testing.T) {

	const BALL_1 = 5
	const BALL_2 = 6
	const BALL_3 = 7
	const TRACK_NAME = "Test"
	const EXPECTED_STATE = TRACK_NAME + ":[5,6,7]"

	// Arrange
	tr := NewTrack(TRACK_NAME, 11, 0)

	// Act
	tr.AddBall(BALL_1)
	tr.AddBall(BALL_2)
	tr.AddBall(BALL_3)

	// Assert
	if tr.String() != EXPECTED_STATE {
		t.Errorf("Unexpected state (actual %v expected, %v)", 
			tr.String(), EXPECTED_STATE)
	}

}