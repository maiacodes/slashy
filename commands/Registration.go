package commands

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
)

func InitialiseCommands() {
	endpoint := fmt.Sprintf("https://discord.com/api/v8/applications/%v/commands", os.Getenv("client_id"))
	if os.Getenv("development") == "true" {
		endpoint = fmt.Sprintf("https://discord.com/api/v8/applications/%v/guilds/%v/commands", os.Getenv("client_id"), os.Getenv("test_server_id"))
	}

	// Fetch existing slash commands and remove ones that aren't routed
	var registeredCommands []commandRoute
	err := fetch(endpoint, "GET", nil, &registeredCommands)
	if err != nil {
		panic(err)
	}

	for _, command := range registeredCommands {
		_, exists := Commands[command.Name]
		if !exists {
			err := fetch(endpoint+"/"+command.ID, "DELETE", nil, nil)
			if err != nil {
				panic(err)
			}
			logrus.Info("Unregistered: ", command.Name)
		}
	}

	// Register commands that aren't already registered
	for _, command := range Commands {
		if !commandExists(registeredCommands, command.Name, command.Description, command.Options) {
			err := fetch(endpoint, "POST", command, nil)
			if err != nil {
				panic(err)
			}
			logrus.Info("Registered: ", command.Name)
		}
	}
}

func commandExists(cmds []commandRoute, nameMatch string, descriptionMath string, optionsMatch []commandOption) bool {
	for _, cmd := range cmds {
		if cmd.Name == nameMatch && cmd.Description == descriptionMath && reflect.DeepEqual(cmd.Options, optionsMatch) {
			return true
		}
	}
	return false
}
