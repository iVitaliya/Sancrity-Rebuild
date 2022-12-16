package utils

type IDummyUser struct {
	bot            bool
	id             string
	tag            string
	name           string
	username       string
	hexAccentColor string
	avatarURL      func() string
}

var (
	DefaultPHP string      = "https://cdn.discordapp.com/embed/avatars/0.png"
	DummyUser  *IDummyUser = &IDummyUser{
		bot:            false,
		id:             "00000000000",
		tag:            "Unknown User#0000",
		name:           "Unknown User",
		username:       "Unknown User",
		hexAccentColor: "#FFFFFF",
		avatarURL:      func() string { return DefaultPHP },
	}
)
