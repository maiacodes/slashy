package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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
	if strings.HasPrefix(URL, "https://discord.com") {
		req.Header.Set("Authorization", "Bot "+os.Getenv("bot_token"))
	}
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
