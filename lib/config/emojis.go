package config

type IEmojiNormal struct {
	Error     string
	Check     string
	Music     string
	Volume    string
	Ball      string
	Christmas string
	Heart     string
	Paper     string
	Scissors  string
	Stone     string
	Pong      string
	Birthday  string
	Clock     string
	Gift      string
	Medal     string
	Party     string
	Info      string
	ArrowDown string
	ArrowUp   string
	Dcredits  string
	Tv        string
	Slash     string
}

type IEmojiBot struct {
	Add  string
	Info string
	Min  string
}

type IEmojiEconomy struct {
	Pocket string
	Bank   string
	Coins  string
}

type IEmojiBadges struct {
	Bot        string
	Management string
	Bug        string
	Developer  string
	Supporter  string
	Team       string
	Booster    string
	Partner    string
	Voter      string
	Support    string
	Moderator  string
	Designer   string
	Active     string
	Event      string
	Vip        string
	Marketing  string
}

type IEmojiAnimated struct {
	Loading string
}

type IEmojis struct {
	Normal   *IEmojiNormal
	Bot      *IEmojiBot
	Economy  *IEmojiEconomy
	Badges   *IEmojiBadges
	Animated *IEmojiAnimated
}

func Emojis() *IEmojis {
	return &IEmojis{
		Normal: &IEmojiNormal{
			Error:     "❌",
			Check:     "✔️",
			Music:     "🎵",
			Volume:    "🔉",
			Ball:      "🎱",
			Christmas: "🎄",
			Heart:     "❤️",
			Paper:     "📰",
			Scissors:  "✂️",
			Stone:     "🪨",
			Pong:      "🏓",
			Birthday:  "🎂",
			Clock:     "⏰",
			Gift:      "🎁",
			Medal:     "🏅",
			Party:     "🎉",
			Info:      "ℹ️",
			ArrowDown: "⬇️",
			ArrowUp:   "⬆️",
			Dcredits:  "💳",
			Tv:        "📺",
			Slash:     "",
		},
		Bot: &IEmojiBot{
			Add:  "📥",
			Info: "ℹ️",
			Min:  "🔻",
		},
		Economy: &IEmojiEconomy{
			Pocket: "👛",
			Bank:   "🏦",
			Coins:  "💰",
		},
		Badges: &IEmojiBadges{
			Bot:        "🤖",
			Management: "👑",
			Bug:        "🐛",
			Developer:  "👨‍💻",
			Supporter:  "👨‍🔧",
			Team:       "👨‍👩‍👧‍👦",
			Booster:    "🚀",
			Partner:    "🤝",
			Voter:      "🗳️",
			Support:    "🔧",
			Moderator:  "👮",
			Designer:   "🎨",
			Active:     "🔥",
			Event:      "🎉",
			Vip:        "👑",
			Marketing:  "📈",
		},
		Animated: &IEmojiAnimated{
			Loading: "🔄️",
		},
	}
}
