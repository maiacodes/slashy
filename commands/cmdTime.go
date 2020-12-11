package commands

import (
	"../interaction"
	"fmt"
	"time"
)

func cmdTime(e *interaction.Event) *interaction.EventCallback {
	location, err := time.LoadLocation(e.Data.Options[0].Value)
	if err != nil {
		return e.Error("Invalid timezone, Make sure to use IANA TZ format!")
	}
	locationTime := time.Now().In(location)
	return e.Reply(fmt.Sprintf("**🕰  |  The time in %v is %v**", location.String(), locationTime.Format("15:04:05 on January 2, 2006")))
}
