package clock

import "fmt"

const testVersion = 4

// Clock create a new struct to manipulate 2 integers as they are hour and minutes
type Clock struct {
	hour   int
	minute int
}

// New return a new Clock struct with given valures of hour and minutes
func New(hour, minute int) Clock {
	return Clock{hour, minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add sum the number in minutes to hour
func (c Clock) Add(minutes int) Clock {
	if minutes < 0 {
		c.minute -= minutes
	} else {
		c.minute += minutes
	}

	return c
}
