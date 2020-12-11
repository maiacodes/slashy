package interaction

type Command struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Options []CommandOption `json:"options"`
}

type CommandOption struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
