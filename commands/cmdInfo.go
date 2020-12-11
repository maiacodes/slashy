package commands

import (
	"../interaction"
)

func cmdInfo(e *interaction.Event) *interaction.EventCallback {
	return e.Reply("Testing bot to test Discord's new slash commands. Operated by Maia")
}
