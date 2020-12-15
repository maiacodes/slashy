package commands

import (
	"github.com/maiacodes/slashy/interaction"
)

func cmdHelp(e *interaction.Event) *interaction.EventCallback {
	return e.ReplyEmbed(interaction.Embed{
		Description: "Here are Slashy's commands!",
		Fields: []interaction.EmbedField{
			interaction.EmbedField{
				Name:  "Check time in location",
				Value: "`/time <location>`",
			},
			interaction.EmbedField{
				Name:  "Convert between currencies",
				Value: "`/currency <from> <amount <to>`",
			},
			interaction.EmbedField{
				Name:  "Evaluate maths",
				Value: "`/maths <expression>`",
			},
			interaction.EmbedField{
				Name:  "Check the responsiveness with a ping",
				Value: "`/ping`",
			},
			interaction.EmbedField{
				Name:  "Make the bot say something",
				Value: "`/say <message>`",
			},
			interaction.EmbedField{
				Name:  "About Slashy",
				Value: "`/info`",
			},
			interaction.EmbedField{
				Name:  "Slashy Commands",
				Value: "`/help`",
			},
		},
		Color: "16762655",
		Author: interaction.EmbedAuthor{
			Name:    "Slashy Help",
			IconURL: "https://media.discordapp.net/attachments/733250810707705877/762359646828363826/swifty.png?width=128&height=128",
			URL:     "https://slashy.maia.codes",
		},
		Footer: interaction.EmbedFooter{
			Text: "Made with ❤️ by Maia & Contributors",
		},
	})
}
