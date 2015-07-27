package control

import (
	"testing"
	//"os"
)

// func TestMain(m *testing.M) {
// 	clockCycles = 0
// 	tracks = []track.Track{}
// 	clockQueue = queue.NewQueue(0)
// 	os.Exit(m.Run())
// }

func TestShouldInitializeClock(t *testing.T) {

	// Arrange
	const EXPECTED_CAPACITY = 30

	// Act
	initClock(EXPECTED_CAPACITY)

	// Assert
	if len(tracks) != 3 {
		t.Errorf("Should return ball 1 [%v]", tracks)
	}

	if getTotalDays() != 0 {
		t.Errorf("Should be no clock cycles [%d]", clockCycles)
	}

}

func TestShouldCalculateTotalDaysByCompletedCycles(t *testing.T) {
	
	// Arrange
	initClock(0)
	clockCycles = 5

	// Act
	totalDays := getTotalDays()

	if totalDays != 2 {
		t.Errorf("Unexpected amount of days (actual %d expected 2)", clockCycles)
	}

}

func TestShouldCalculateHalfDaysIntoTotalMinutes(t *testing.T) {
	
	// Arrange
	initClock(0)
	clockCycles = 5
	// Act
	totalMinutes := getTotalMinutes()

	if totalMinutes != 3600 {
		t.Errorf("Unexpected amount of minutes (actual %d expected 3600) [clockCycles:%d, tracks:%v]", 
			totalMinutes, clockCycles, tracks)
	}

}

func TestShouldCalculateCycledMinutesAndCurrentStateMinutes(t *testing.T) {
	
	// Arrange
	initClock(0)
	clockCycles = 5

	// Add balls to hour track, 120 minutes
	tracks[2].AddBall(1)
	tracks[2].AddBall(2)

	// Add balls to 5 minute track, 15 minutes
	tracks[1].AddBall(3)
	tracks[1].AddBall(4)
	tracks[1].AddBall(5)

	// Add balls to minute track, 4 minutes
	tracks[0].AddBall(6)
	tracks[0].AddBall(7)
	tracks[0].AddBall(8)
	tracks[0].AddBall(9)

	// Act
	totalMinutes := getTotalMinutes()

	if totalMinutes != 3600 + 120 + 15 + 4 {
		t.Errorf("Unexpected amount of minutes (actual %d expected 3739) [clockCycles:%d, tracks:%v]", 
			totalMinutes, clockCycles, tracks)
	}

}

func TestShouldDigest1BallAndReturn1Minute(t *testing.T) {

	// Arrange
	initClock(10)

	// Act
	digestBall(1)

	totalMinutes := getTotalMinutes()

	if totalMinutes != 1 {
		t.Errorf("Unexpected amount of minutes (actual %d expected 1) [clockCycles:%d, tracks:%v]", 
			totalMinutes, clockCycles, tracks)
	}

}

func TestShouldDigest2BallsAndReturn2Minutes(t *testing.T) {

	// Arrange
	initClock(10)

	// Act
	digestBall(1)
	digestBall(2)

	totalMinutes := getTotalMinutes()

	if totalMinutes != 2 {
		t.Errorf("Unexpected amount of minutes (actual %d expected 2) [clockCycles:%d, tracks:%v]", 
			totalMinutes, clockCycles, tracks)
	}

}

func TestShouldDigest5BallsAndReturn5Minutes(t *testing.T) {

	// Arrange
	initClock(10)

	// Act
	digestBall(1)
	digestBall(2)
	digestBall(3)
	digestBall(4)
	digestBall(5)

	totalMinutes := getTotalMinutes()

	if totalMinutes != 5 {
		t.Errorf("Unexpected amount of minutes (actual %d expected 5) [clockCycles:%d, tracks:%v]", 
			totalMinutes, clockCycles, tracks)
	}

}

func TestShouldDigest1HourOfBalls(t *testing.T) {

	// Arrange
	initClock(100)
	ballCount := 60

	// Act
	for i := 1; i <= ballCount; i++ {
		digestBall(i)
	}
	
	totalMinutes := getTotalMinutes()

	if totalMinutes != 60 {
		t.Errorf("Unexpected amount of minutes (actual %d expected 60) [clockCycles:%d, tracks:%v]", 
			totalMinutes, clockCycles, tracks)
	}

}

func TestShouldDigest2HoursOfBalls(t *testing.T) {

	// Arrange
	initClock(120)
	ballCount := 120

	// Act
	for i := 1; i <= ballCount; i++ {
		digestBall(i)
	}
	
	totalMinutes := getTotalMinutes()

	if totalMinutes != 120 {
		t.Errorf("Unexpected amount of minutes (actual %d expected 120) [clockCycles:%d, tracks:%v]", 
			totalMinutes, clockCycles, tracks)
	}

}

func TestShouldDigestFullClockCycle(t *testing.T) {

	// Arrange
	initClock(720)
	ballCount := 720

	// Act
	for i := 1; i <= ballCount; i++ {
		digestBall(i)
	}
	
	totalMinutes := getTotalMinutes()

	if totalMinutes != 720 {
		t.Errorf("Unexpected amount of minutes (actual %d expected 720) [clockCycles:%d, tracks:%v]", 
			totalMinutes, clockCycles, tracks)
	}

	if clockCycles != 1 {
		t.Errorf("Unexpected clock cycle tick (actual %d expected 1) [tracks:%v]", 
			clockCycles, tracks)
	}

}

func TestShouldDigestAlmost2Cycles(t *testing.T) {

	// Arrange
	initClock(1440)
	ballCount := 1439

	// Act
	for i := 1; i <= ballCount; i++ {
		digestBall(i)
	}
	
	totalMinutes := getTotalMinutes()

	if totalMinutes != 1439 {
		t.Errorf("Unexpected amount of minutes (actual %d expected 1439) [clockCycles:%d, tracks:%v]", 
			totalMinutes, clockCycles, tracks)
	}

	if clockCycles != 1 {
		t.Errorf("Unexpected clock cycle tick (actual %d expected 1) [tracks:%v]", 
			clockCycles, tracks)
	}

}

func TestShouldReturnError(t *testing.T) {

	// Arrange
	const CAPACITY = 0

	// Act
	_, err := Run(CAPACITY, 0)
	
	if err == nil {
		t.Errorf("Expected Error")
	}

}

func TestShouldEndWhenReachedDesiredMinutes(t *testing.T) {

	// Arrange
	const CAPACITY = 30

	// Act
	Run(CAPACITY, 325)
	
	if getTotalMinutes() != 325 {
		t.Errorf("Unexpected amount of minutes (actual %d expected 325) [clockCycles:%d, tracks:%v, queue:%v]", 
			getTotalMinutes(), clockCycles, tracks, clockQueue)
	}

}

func TestShouldRunFor15Days(t *testing.T) {

	// Arrange
	const CAPACITY = 30

	// Act
	totalDays, _ := Run(CAPACITY, 0)
	
	if totalDays != 15 {
		t.Errorf("Unexpected total days (actual %d expected 15) [minutes:%v, clockCycles:%d, tracks:%v, queue:%v]", 
			totalDays, getTotalMinutes(), clockCycles, tracks, clockQueue)
	}

}

func TestShouldRunFor378Days(t *testing.T) {

	// Arrange
	const CAPACITY = 45

	// Act
	totalDays, _ := Run(CAPACITY, 0)
	
	if totalDays != 378 {
		t.Errorf("Unexpected total days (actual %d expected 15) [minutes:%v, clockCycles:%d, tracks:%v, queue:%v]", 
			totalDays, getTotalMinutes(), clockCycles, tracks, clockQueue)
	}

}