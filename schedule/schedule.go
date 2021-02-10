package schedule

import (
	"container/heap"
	"time"
)

// Schedule is the base struct representing a schedule
type Schedule struct {
	Schedule []time.Time `json:"schedule"`
}

// Assert that Schedule implements the heap interface at compile time
var _ heap.Interface = &Schedule{}

// Implementing the heap interface
func (s Schedule) Len() int           { return len(s.Schedule) }
func (s Schedule) Less(i, j int) bool { return s.Schedule[i].Sub(s.Schedule[j]) < time.Duration(0) }
func (s Schedule) Swap(i, j int)      { s.Schedule[i], s.Schedule[j] = s.Schedule[j], s.Schedule[i] }

// Push is used make Schedule implement the heap interface
func (s *Schedule) Push(x interface{}) {
	s.Schedule = append(s.Schedule, x.(time.Time))
}

// Pop is used make Schedule implement the heap interface
func (s *Schedule) Pop() interface{} {
	old := s.Schedule
	n := len(old)
	x := old[n-1]
	s.Schedule = old[0 : n-1]
	return x
}

// NewSchedule creates a Schedule object
// It will error if the required fields are improperly filled in.
func NewSchedule(layout, value, location string) (*Schedule, error) {
	s := &Schedule{}
	err := s.SetTime(layout, value, location)

	if err != nil {
		return nil, err
	}
	return s, nil
}

// SetTime parses a string input and sets the time
// for the event in the Schedule.
func (s *Schedule) SetTime(layout, value, location string) error {
	// if location == "" {
	// 	location = "Local"
	// }
	loc, err := time.LoadLocation(location)
	if err != nil {
		return err
	}

	t, err := parse(layout, value, loc)
	if err != nil {
		return err
	}
	s.CreateSchedule(t)
	return nil
}

// parse wraps time.Parse to ensure that the local time is
// returned if a duration is not set.
func parse(layout, value string, location *time.Location) (time.Time, error) {
	if location == nil {
		return time.Parse(layout, value)
	}
	return time.ParseInLocation(layout, value, location)
}

// CreateSchedule creates a Schedule
func (s *Schedule) CreateSchedule(t time.Time) {
	s.Schedule = []time.Time{t}
	heap.Init(s)
}

// Next returns the next scheduled element
func (s *Schedule) Next() time.Time {
	t := heap.Pop(s)
	return t.(time.Time)
}
