package schedule

import (
	"container/heap"
	"errors"
	"fmt"
	"time"
)

// Repeater is an interface whose implementations
// define how to repeat a schedule.
// For example:
// a) Repeating at fixed intervals
// b) Repeating on certain days of the week
type Repeater interface {
	Repeat(schedule *Schedule) error
}

// Ensure at compile time that the Repeater Structs implement the Repeater interface
var _ Repeater = &IntervalRepeater{}
var _ Repeater = &DayOfWeekRepeater{}

// IntervalRepeater repeats the
// schedule NumTimes with time between repeats of
// Interval. NumTimes does not include the original
// schedule.
type IntervalRepeater struct {
	NumTimes int
	Interval *time.Duration
}

// Repeat takes a Schedule object and alters it
// to repeat ir.NumTimes times at fixed intervals.
func (ir *IntervalRepeater) Repeat(s *Schedule) error {
	if ir.NumTimes < 1 {
		return fmt.Errorf(
			"repeater: NumTimes is set to %d - it should be at least 1",
			ir.NumTimes,
		)
	}
	if ir.Interval == nil {
		return errors.New("repeater: Interval is set to nil")
	}

	old := s.Schedule
	n := len(old)
	new := make([]time.Time, n*(ir.NumTimes+1))
	interval := *ir.Interval
	copy(new, old)
	for i := n; i < n*(ir.NumTimes+1); i++ {
		new[i] = new[i-n].Add(interval)
	}
	s.Schedule = new
	heap.Init(s)
	return nil
}

// DayOfWeekRepeater repeats the schedule NumTimes
// at the same time of day, but only on specified
// days of the week. NumTimes does not include the
// original schedule.
type DayOfWeekRepeater struct {
	NumTimes int
	Days     []time.Weekday
}

// Repeat takes a Schedule and alters it
// to repeat ir.NumTimes times, repeating
// only on certain weekdays.
func (r *DayOfWeekRepeater) Repeat(s *Schedule) error {
	if r.NumTimes < 1 {
		return fmt.Errorf(
			"repeater: NumTimes is set to %d - it should be at least 1",
			r.NumTimes,
		)
	}
	if len(r.Days) == 0 {
		return errors.New("repeater: No Days Of Week provided to repeat")
	}

	old := s.Schedule
	n := len(old)
	new := make([]time.Time, n*(r.NumTimes+1))
	copy(new, old)

	day := time.Duration(24) * time.Hour

	wdayMap := r.getWeekdayMap()
	for i := n; i < n*(r.NumTimes+1); i++ {
		t := new[i-n].Add(day)
		wday := int(t.Weekday())
		new[i] = t.Add(wdayMap[wday])
	}
	s.Schedule = new
	heap.Init(s)
	return nil
}

func (r *DayOfWeekRepeater) getWeekdayMap() map[int]time.Duration {
	wdays := [7]int{}
	for _, v := range r.Days {
		wdays[v] = 1
	}
	wdayMap := make(map[int]time.Duration)
	for i, v := range wdays {
		if v == 1 {
			wdayMap[i] = 0
		} else {
			j := 1
			for wdays[(i+j)%7] == 0 {
				j++
			}
			wdayMap[i] = time.Duration(j*24) * time.Hour
		}
	}
	return wdayMap
}
