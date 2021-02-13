package interaction

import "time"

// Timestamp is an ISO8601 format time string
// e.g. 2021-02-13T01:45:21+0000
type Timestamp string

func (t Timestamp) Time() (tm time.Time, err error) {
	// RFC3339 and ISO8601 are almost the same thing
	return time.Parse(time.RFC3339, string(t))
}
