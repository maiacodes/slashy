package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
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

func fetch(URL string, Method string, RequestBody interface{}, ResponseBody interface{}) (err error) {
	//Make body
	body := []byte("")
	if RequestBody != nil {
		body, _ = json.Marshal(RequestBody)
	}
	req, _ := http.NewRequest(Method, URL, bytes.NewReader(body))
	if RequestBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", "Bot "+os.Getenv("bot_token"))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	//Format response body
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 204 {
		err = fmt.Errorf(resp.Status + ": " + string(body))
		return
	}

	defer resp.Body.Close()

	if ResponseBody != nil {
		//Format JSON into an object Struct
		err = json.Unmarshal(body, &ResponseBody)
		if err != nil {
			return
		}
	}

	return
}
