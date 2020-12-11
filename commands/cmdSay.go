package commands

import (
	"../interaction"
)

func cmdSay(e *interaction.Event) *interaction.EventCallback {
	return e.Reply(e.Data.Options[0].Value)
}
