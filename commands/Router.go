package commands

import (
	"../interaction"
)

func Router(e *interaction.Event) *interaction.EventCallback {
	route, ok := Commands[e.Data.Name]
	if !ok {
		return e.Reply("**❌  |  Command not found.**")
	}
	return route.CommandFunction(e)
}
