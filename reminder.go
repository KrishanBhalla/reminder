package reminder

import (
	"github.com/KrishanBhalla/reminder/notify"
	"github.com/KrishanBhalla/reminder/schedule"
)

// Reminder implements a scheduled reminder
type Reminder struct {
	Title    string
	Message  string
	Schedule *schedule.Schedule
	Notifier notify.Notifier
}

// Notify sends the notification
func (r *Reminder) Notify() error {
	return r.Notifier.Notify(r.Title, r.Message)
}
