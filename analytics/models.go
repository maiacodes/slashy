package analytics

import "time"

type event struct {
	GuildID   string
	EventType string
	Timestamp time.Time
}
