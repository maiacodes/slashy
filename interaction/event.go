package interaction

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Event struct {
	ChannelID string  `json:"channel_id"`
	Token     string  `json:"token"`
	Type      int     `json:"type"`
	Version   int     `json:"version"`
	GuildID   string  `json:"guild_id"`
	Data      Command `json:"data"`
	ID        string  `json:"id"`
	Responded bool
}

func (e Event) Reply(content string) *EventCallback {
	callback := EventCallback{
		Type: CallbackWithSource,
		Data: EventCallbackData{Content: content},
	}

	// If the request hasn't been responded yet, send the callback
	if !e.Responded {
		return &callback
	}

	// If the request has been responded already, send the callback manually
	body, _ := json.Marshal(callback)
	req, _ := http.NewRequest("POST", fmt.Sprintf("https://discord.com/api/v8/interactions/%v/%v/callback", e.ID, e.Token), bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	client.Do(req)
	// TODO: error handling
	return nil
}
