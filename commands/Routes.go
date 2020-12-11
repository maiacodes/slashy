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

const (
	OPTION_TYPE_STRING  = 3
	OPTION_TYPE_INT     = 4
	OPTION_TYPE_BOOL    = 5
	OPTION_TYPE_USER    = 6
	OPTION_TYPE_CHANNEL = 7
	OPTION_TYPE_ROLE    = 8
)

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
					Type:        OPTION_TYPE_STRING,
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
		"maths": {
			Name:            "maths",
			Description:     "Evaluate a maths expression.",
			CommandFunction: cmdMaths,
			Options: []commandOption{
				{
					Type:        OPTION_TYPE_STRING,
					Name:        "expression",
					Description: "What you'd like to evaluate.",
					Required:    true,
				},
			},
		},
		"currency": {
			Name:            "currency",
			Description:     "Check exchange rates.",
			CommandFunction: cmdCurrency,
			Options: []commandOption{
				{
					Type:        OPTION_TYPE_STRING,
					Name:        "from",
					Description: "What currency you're converting from.",
					Required:    true,
				},
				{
					Type:        OPTION_TYPE_STRING,
					Name:        "amount",
					Description: "The amount of the base currency you're converting.",
					Required:    true,
				},
				{
					Type:        OPTION_TYPE_STRING,
					Name:        "to",
					Description: "What currency you're converting into.",
					Required:    true,
				},
			},
		},
	}
)
