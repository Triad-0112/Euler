package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	Minutes int
}

const (
	daytmin = 1440
)

func New(h, m int) Clock {
	return Clock{((h*60+m)%daytmin + daytmin) % daytmin}
}

func (c Clock) Add(m int) Clock {
	return New(0, c.Minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.Minutes-m)
}

func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Minutes/60, c.Minutes%60)
}
