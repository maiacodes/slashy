package commands

import (
	"github.com/maiacodes/slashy/interaction"
)

func cmdSay(e *interaction.Event) *interaction.EventCallback {
	return e.Reply(e.Data.Options[0].Value)
}
