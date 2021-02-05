package analytics

import (
	"github.com/sirupsen/logrus"
	"time"
)

func Track(GuildID string, EventType string) {
	if G == nil {
		return
	}

	err := G.Create(&event{
		GuildID:   GuildID,
		EventType: EventType,
		Timestamp: time.Now(),
	}).Error

	if err != nil {
		logrus.Error(err)
	}
}
