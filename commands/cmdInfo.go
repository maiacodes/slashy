package commands

import (
	"../interaction"
)

func cmdInfo(e *interaction.Event) *interaction.EventCallback {
	return e.ReplyEmbed(interaction.Embed{
		Description: "Slashy is an [open-source](https://github.com/maiacodes/slashy) slash commands bot written in Golang!",
		Color:       "16762655",
		Author: interaction.EmbedAuthor{
			Name:    "About Slashy",
			IconURL: "https://media.discordapp.net/attachments/733250810707705877/762359646828363826/swifty.png?width=128&height=128",
			URL:     "https://slashy.maia.codes",
		},
		Footer: interaction.EmbedFooter{
			Text: "Made with ❤️ by Maia & Contributors",
		},
	})
}
