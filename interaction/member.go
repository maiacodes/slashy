package interaction

type Member struct {
	User         User        `json:"user"`
	Nick         string      `json:"nick"`
	Roles        []string    `json:"roles"`         // array of roles IDs
	JoinedAt     Timestamp   `json:"joined_at"`     // the time since the user joined this guild
	PremiumSince Timestamp   `json:"premium_since"` // the time since the user had discord nitro
	Deafened     bool        `json:"deaf"`
	Muted        bool        `json:"mute"`
	Permissions  Permissions `json:"permissions"`
}

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
	SystemUser    bool   `json:"system"`
	MFAEnabled    bool   `json:"mfa_enabled"`
	Locale        string `json:"locale"`
}
