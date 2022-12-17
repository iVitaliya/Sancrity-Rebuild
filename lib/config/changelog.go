package config

type IChangelog struct {
	Release  string
	Bugs     []string
	Features []string
}

var Changelog *IChangelog = &IChangelog{
	Release: "16-12-2022",
	Bugs: []string{
		"Better bot setups",
		"Giveaways fixed",
		"A new invite tracker system",
		"Channel logs bug fixed",
		"Better error handling for commands",
	},
	Features: []string{
		"Bot completely in slash commands",
		"New Activities",
		"New server stats",
		"Auto setups",
		"Scheduled Events logs",
		"Mute command is now timeout command",
		"Embed layout improvements",
		"A advanced embed builder",
		"Removed snipe commands",
		"Unnecessary commands removed",
		"Some commands rearranged",
		"**Bot 1 year pack** available for **10 Dcredits**",
		"Better error handling for commands",
	},
}
