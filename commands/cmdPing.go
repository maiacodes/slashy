package commands

import (
	"../interaction"
)

func cmdPing(e *interaction.Event) *interaction.EventCallback {
	return e.Reply("**🏓  |  Pong!**")
}
