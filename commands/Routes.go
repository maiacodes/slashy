package commands

import (
	"../interaction"
)

type commandRoute struct {
	ID              string                                                `json:"id"`
	Name            string                                                `json:"name"`
	Description     string                                                `json:"description"`
	CommandFunction func(e *interaction.Event) *interaction.EventCallback `json:"-"`
	Options         []commandOption                                       `json:"options"`
}

type commandOption struct {
	Type        int    `json:"type"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
}

var (
	Commands = map[string]commandRoute{
		"ping": {
			Name:            "ping",
			Description:     "Check the responsiveness of the bot.",
			CommandFunction: cmdPing,
		},
		"say": {
			Name:            "say",
			Description:     "Make the bot say something!",
			CommandFunction: cmdSay,
			Options: []commandOption{
				{
					Type:        3,
					Name:        "message",
					Description: "What you'd like the bot to say.",
					Required:    true,
				},
			},
		},
		"info": {
			Name:            "info",
			Description:     "About the bot.",
			CommandFunction: cmdInfo,
		},
	}
)
