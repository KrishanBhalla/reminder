# reminder
Reminder is a Golang package to allow users to schedule alerts.
It has 4 parts:

1. Scheduler
2. Repeater
3. Notifier
4. Reminder
 
A scheduler takes in a `time.Time` for the first reminder to be sent.

A repeater applies rules to repeating a scheduled event - say at fixed intervals or on certain days
Repeaters can be combined (see examples)

The notifier defines how one gets notified. By default the `beeep` package is used to provide desktop notifications

The reminder pulls it all together, and allows the user to send a specified message on a given schedule.
## Examples
### Simple
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

    format := time.RFC1123Z
    s, _ := schedule.NewSchedule(format, time.Now().Format(format), "GMT")

    r := schedule.IntervalRepeater{
        Interval: time.Duration(5) * time.Second,
        NumTimes: 3,
    }
    r.Repeat(s)

    rem := reminder.Reminder{
        Schedule: s,
        Notifier: &notify.DesktopNotifier{},
    }

    err := rem.Remind("Reminder", "I'm a reminder made in Go!")
    fmt.Println(err)
}
```
### Combining Repeaters
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

	format := time.RFC1123Z
	s, _ := schedule.NewSchedule(format, time.Now().Format(format), "GMT")

	r := schedule.IntervalRepeater{
		Interval: time.Hour,
		NumTimes: 3,
	}
	r.Repeat(s)

    wdayRem := schedule.DayOfWeekRepeater{
		Interval: time.Hour,
		NumTimes: 5,
	}

	rem := reminder.Reminder{
		Schedule: s,
		Notifier: &notify.DesktopNotifier{},
	}

	err := rem.Remind("Excercise", "Leave the desk and strech your legs")
	fmt.Println(err)

}

```
