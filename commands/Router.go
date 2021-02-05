package commands

import (
	"github.com/maiacodes/slashy/analytics"
	"github.com/maiacodes/slashy/interaction"
)

func Router(e *interaction.Event) *interaction.EventCallback {
	route, ok := Commands[e.Data.Name]
	if !ok {
		return e.Reply("**❌  |  Command not found.**")
	}
	go analytics.Track(e.GuildID, route.Name)
	return route.CommandFunction(e)
}
