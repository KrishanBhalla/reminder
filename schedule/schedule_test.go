package schedule

import (
	"testing"
	"time"
)

func TestScheduleSetTime(t *testing.T) {
	// arrange
	schedule := &Schedule{}
	locations := []string{"Local", "Europe/London", "America/New_York", ""}
	layout := time.RFC1123Z
	value := time.Now().Format(layout)

	for _, loc := range locations {
		// act
		err := schedule.SetTime(layout, value, loc)
		// assert
		if err != nil {
			t.Errorf(
				"Failed to set time with location: %s \n. Received error: %s",
				loc,
				err.Error(),
			)
		}
	}

}
