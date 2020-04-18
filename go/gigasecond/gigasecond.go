// Package gigasecond contains methods to add a gigasecond to a time.
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond returns the time input with a gigasecond added to it.
func AddGigasecond(t time.Time) time.Time {

	return t.Add(time.Second * 1000000000)
}
