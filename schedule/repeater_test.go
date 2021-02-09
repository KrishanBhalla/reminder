package schedule

import (
	"testing"
	"time"
)

func testingIntervalRepeater(n int, interval time.Duration) (s *Schedule, r Repeater) {
	layout := time.RFC1123Z
	value := time.Now().Format(layout)

	schedule, _ := NewSchedule(layout, value, "Local")

	repeater := &IntervalRepeater{
		NumTimes: n,
		Interval: interval,
	}
	return schedule, repeater
}

func testingDayOfWeekRepeater(n int, weekdays ...time.Weekday) (s *Schedule, r Repeater) {
	layout := time.RFC1123Z
	value := time.Now().Format(layout)

	schedule, _ := NewSchedule(layout, value, "Local")

	repeater := &DayOfWeekRepeater{
		NumTimes: n,
		Days:     weekdays,
	}
	return schedule, repeater
}

func TestIntervalRepeaterCanRepeat(t *testing.T) {
	// arrange
	n := 10
	interval := time.Duration(5) * time.Minute
	s, r := testingIntervalRepeater(n, interval)
	// act
	m := s.Len()
	r.Repeat(s)
	// assert
	if s.Len() != m*(n+1) {
		t.Errorf(
			"Repeated schedule did not have expected length. Expected %d, returned %d",
			m*(n+1),
			s.Len(),
		)
	}

	// Test intervals
	pres := s.Next()
	for s.Len() > 0 {
		next := s.Next()
		if d := next.Sub(pres); d != interval {
			t.Errorf(
				"Consecutive scheduled events did not have expected interval. Expected %s, Retured %s",
				interval,
				d,
			)
		}
		pres = next
	}
}

func TestIntervalRepeaterCannotRepeatWithZeroDuration(t *testing.T) {
	// arrange
	n := 10
	interval := time.Duration(0)
	s, r := testingIntervalRepeater(n, interval)
	// act
	err := r.Repeat(s)
	// assert
	if err == nil {
		t.Errorf("Expected error when passing in 0 duration to repeater")
	}
}

func TestIntervalRepeaterCannotRepeatWithZeroTimes(t *testing.T) {
	// arrange
	n := 0
	interval := time.Duration(5) * time.Minute
	s, r := testingIntervalRepeater(n, interval)
	// act
	err := r.Repeat(s)
	// assert
	if err == nil {
		t.Errorf("Expected error when passing in 0 times to repeater")
	}
}

func TestDayOfWeekRepeaterCanRepeat(t *testing.T) {
	// arrange
	n := 10
	day := time.Monday
	s, r := testingDayOfWeekRepeater(n, day)
	// act
	m := s.Len()
	r.Repeat(s)
	// assert
	if s.Len() != m*(n+1) {
		t.Errorf(
			"Repeated schedule did not have expected length. Expected %d, returned %d",
			m*(n+1),
			s.Len(),
		)
	}

	// Test intervals
	i := 0
	for s.Len() > 0 {
		pres := s.Next()
		if i < m {
			i++
			continue
		}
		if wd := pres.Weekday(); wd != day {
			t.Errorf(
				"Event was not on expected weekday. Expected %s, Retured %s",
				day,
				wd,
			)
		}
		pres = s.Next()
	}
}

func TestDayOfWeekCannotRepeatWithNoDays(t *testing.T) {
	// arrange
	n := 10
	s, r := testingDayOfWeekRepeater(n)
	// act
	err := r.Repeat(s)
	// assert
	if err == nil {
		t.Errorf("Expected error when passing in 0 duration to repeater")
	}
}

func TestDayOfWeekcannotRepeatWithZeroTimes(t *testing.T) {
	// arrange
	n := 0
	days := []time.Weekday{time.Monday, time.Tuesday}
	s, r := testingDayOfWeekRepeater(n, days...)
	// act
	err := r.Repeat(s)
	// assert
	if err == nil {
		t.Errorf("Expected error when passing in 0 times to repeater")
	}
}
