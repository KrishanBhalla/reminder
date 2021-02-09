package notify

import "github.com/gen2brain/beeep"

// Notifier is a generic interface for implementing a notification
// TODO: Consider updating to use github.com/nikoksr/notify?
type Notifier interface {
	Notify(title, message string) error
}

// DesktopNotifier wraps  github.com/gen2brain/beeep
type DesktopNotifier struct {
}

// Notify creates a desktop notification
func (d *DesktopNotifier) Notify(title, message string) error {
	return beeep.Notify(title, message, "assets/reminder.png")
}
