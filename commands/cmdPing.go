package commands

import (
	"github.com/maiacodes/slashy/interaction"
)

func cmdPing(e *interaction.Event) *interaction.EventCallback {
	return e.Reply("**🏓  |  Pong!**")
}
