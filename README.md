[![Codacy Badge](https://app.codacy.com/project/badge/Grade/e6eb25534d2c404f8ce31ca3b4dbfa63)](https://www.codacy.com/gh/KrishanBhalla/reminder/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=KrishanBhalla/reminder&amp;utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com//KrishanBhalla/reminder)](https://goreportcard.com/report/github.com/KrishanBhalla/reminder)

# reminder

Reminder is a Golang package to allow users to schedule alerts.
It has 4 parts:

1.  Scheduler
2.  Repeater
3.  Notifier
4.  Reminder
 
A scheduler takes in a `time.Time` for the first reminder to be sent.

A repeater applies rules to repeating a scheduled event - say at fixed intervals or on certain days

Repeaters can be combined (see examples)

The notifier defines how one gets notified. By default the `beeep` package is used to provide desktop notifications

The reminder pulls it all together, and allows the user to send a specified message on a given schedule.
## Examples
### Schedule a single event
```go
package main

import (
	"fmt"
	"time"

	"github.com/KrishanBhalla/reminder"
	"github.com/KrishanBhalla/reminder/notify"
	"github.com/KrishanBhalla/reminder/schedule"
)

func main() {
	// Schedule an event
    format := time.RFC1123Z
	t := time.Now().Add(time.Minute)
    s, _ := schedule.NewSchedule(format, t.Format(format), "Local")
	// Create a reminder
    rem := reminder.Reminder{
        Schedule: s,
        Notifier: &notify.Desktop{},
    }

    err := rem.Remind("Reminder", "I'm a reminder made in Go!")
    fmt.Println(err)
}
```
### One Repeater
```go
package main

import (
	"fmt"
	"time"

	"github.com/KrishanBhalla/reminder"
	"github.com/KrishanBhalla/reminder/notify"
	"github.com/KrishanBhalla/reminder/schedule"
)

func main() {

	// Schedule an event
    format := time.RFC1123Z
    s, _ := schedule.NewSchedule(format, time.Now().Format(format), "Local")

	// Repeat it every 5 seconds
    r := schedule.IntervalRepeater{
        Interval: time.Duration(5) * time.Second,
        NumTimes: 3,
    }
    r.Repeat(s)
	// Create the reminder
    rem := reminder.Reminder{
        Schedule: s,
        Notifier: &notify.Desktop{},
    }

    err := rem.Remind("Reminder", "I'm a reminder made in Go!")
    fmt.Println(err)
}
```
### Combining Repeaters and sending to Pushbullet
```go

package main

import (
	"fmt"
	"time"

	"github.com/KrishanBhalla/reminder"
	"github.com/KrishanBhalla/reminder/notify"
	"github.com/KrishanBhalla/reminder/schedule"
)

func main() {

	// Schedule a reminder for right now
	format := time.RFC1123Z
	s, _ := schedule.NewSchedule(format, time.Now().Format(format), "Local")

	// Repeat it every hour for the next 8 hours
	r := schedule.IntervalRepeater{
		Interval: time.Hour,
		NumTimes: 8,
	}
	r.Repeat(s)

	// Repeat all of those reminders every day for a week
	wdayRem := schedule.DayOfWeekRepeater{
		Days:     []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday},
		NumTimes: 5,
	}
	wdayRem.Repeat(s)

	notifier := notify.NewPushbullet("apiToken", "Device1", "Device2")
	rem := reminder.Reminder{
		Schedule: s,
		Notifier: notifier,
	}

	err := rem.Remind("Exercise", "Leave the desk and stretch your legs")
	fmt.Println(err)

}


```
## Author
[@KrishanBhalla](https://github.com/KrishanBhalla/)
## Road Map
1.  More thorough unit testing.
2.  Concurrency and simplification of setting up indepenedent reminders.
3.  Other notification services
4.  UI
## Credits
[beeep](https://github.com/gen2brain/beeep)
