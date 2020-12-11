package interaction

type Embed struct {
	Description string       `json:"description"` // The main content box of the embed
	Color       string       `json:"color"`       // The embed's color in int hex
	Fields      []EmbedField `json:"fields"`
	Image       EmbedImage   `json:"image"`  // Image
	Author      EmbedAuthor  `json:"author"` // Author (top)
	Footer      EmbedFooter  `json:"footer"` // Footer (bottom)
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type EmbedAuthor struct {
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
	URL     string `json:"url"`
}

type EmbedFooter struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}

type EmbedImage struct {
	URL string `json:"url"`
}
