package reminder

import (
	"fmt"
	"time"

	"github.com/KrishanBhalla/reminder/notify"
	"github.com/KrishanBhalla/reminder/schedule"
)

// Reminder implements a scheduled reminder
type Reminder struct {
	Schedule *schedule.Schedule
	Notifier notify.Notifier
}

// Remind will loop through all scheduled items, and
// notify the user
func (r *Reminder) Remind(title, message string) error {

	s := *(r.Schedule)
	for s.Len() > 0 {
		s := s
		if d := time.Until(s.Next()); d > time.Duration(0) {
			time.Sleep(d)
		}
		fmt.Println("In loop")
		err := r.Notifier.Notify(title, message)
		if err != nil {
			return err
		}
	}
	return nil
}
