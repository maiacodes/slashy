package commands

import (
	"fmt"
	"github.com/maiacodes/slashy/interaction"
)

type query struct {
	Word     string    `json:"word"`
	Meanings []meaning `json:"meanings"`
}

type meaning struct {
	Class       string       `json:"partOfSpeech"`
	Definitions []definition `json:"definitions"`
}

type definition struct {
	Definition string `json:"definition"`
	Example    string `json:"example"`
}

func cmdDictionary(e *interaction.Event) *interaction.EventCallback {
	def := e.Data.Options[0].Value

	// Fetch
	var resp []query
	err := fetch("https://api.dictionaryapi.dev/api/v2/entries/en_US/"+def, "GET", nil, &resp)
	if err != nil {
		return e.Error(err.Error())
	}

	if len(resp) == 0 || len(resp[0].Meanings) == 0 {
		return e.Error("We cannot find a definition for this word.")
	}

	// Construct embed
	emb := interaction.Embed{
		Color:       "16762655",
		Description: "Definitions of " + def,
		Author: interaction.EmbedAuthor{
			Name:    "Slashy Dictionary",
			IconURL: "https://media.discordapp.net/attachments/733250810707705877/762359646828363826/swifty.png?width=128&height=128",
			URL:     "https://slashy.maia.codes",
		},
		Footer: interaction.EmbedFooter{
			Text: "Data from DictionaryAPI.dev",
		},
	}

	for _, c := range resp {
		for _, d := range c.Meanings {
			for _, b := range d.Definitions {
				emb.Fields = append(emb.Fields, interaction.EmbedField{
					Name: fmt.Sprintf("[__%v__] %v", d.Class, b.Definition),
					Value: fmt.Sprintf("*%v*", func(s string) string {
						if s == "" {
							return "No examples"
						} else {
							return s
						}
					}(b.Example)),
					Inline: false,
				})
			}
		}
	}

	return e.ReplyEmbed(emb)
}
