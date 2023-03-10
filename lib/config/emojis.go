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
			Error:     "โ",
			Check:     "โ๏ธ",
			Music:     "๐ต",
			Volume:    "๐",
			Ball:      "๐ฑ",
			Christmas: "๐",
			Heart:     "โค๏ธ",
			Paper:     "๐ฐ",
			Scissors:  "โ๏ธ",
			Stone:     "๐ชจ",
			Pong:      "๐",
			Birthday:  "๐",
			Clock:     "โฐ",
			Gift:      "๐",
			Medal:     "๐",
			Party:     "๐",
			Info:      "โน๏ธ",
			ArrowDown: "โฌ๏ธ",
			ArrowUp:   "โฌ๏ธ",
			Dcredits:  "๐ณ",
			Tv:        "๐บ",
			Slash:     "",
		},
		Bot: &IEmojiBot{
			Add:  "๐ฅ",
			Info: "โน๏ธ",
			Min:  "๐ป",
		},
		Economy: &IEmojiEconomy{
			Pocket: "๐",
			Bank:   "๐ฆ",
			Coins:  "๐ฐ",
		},
		Badges: &IEmojiBadges{
			Bot:        "๐ค",
			Management: "๐",
			Bug:        "๐",
			Developer:  "๐จโ๐ป",
			Supporter:  "๐จโ๐ง",
			Team:       "๐จโ๐ฉโ๐งโ๐ฆ",
			Booster:    "๐",
			Partner:    "๐ค",
			Voter:      "๐ณ๏ธ",
			Support:    "๐ง",
			Moderator:  "๐ฎ",
			Designer:   "๐จ",
			Active:     "๐ฅ",
			Event:      "๐",
			Vip:        "๐",
			Marketing:  "๐",
		},
		Animated: &IEmojiAnimated{
			Loading: "๐๏ธ",
		},
	}
}
