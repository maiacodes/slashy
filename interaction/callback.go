package interaction

type EventCallbackType int

const (
	CallbackNoSource   = 3
	CallbackWithSource = 4
)

type EventCallback struct {
	Type EventCallbackType `json:"type"`
	Data EventCallbackData `json:"data"`
}

type EventCallbackData struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds"`
}
