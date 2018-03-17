package configuration

// CurrentConfig will be set after loading so it can be accessed from outside
var CurrentConfig Config

// Config contains the bot config parameters
type Config struct {
	Token         string        `json:"token"`
	CommandPrefix string        `json:"commandPrefix"`
	Commands      CommandConfig `json:"commands"`
}

// CommandConfig holds the information about the bot commands
type CommandConfig struct {
	Price   CommandPriceConfig   `json:"price"`
	Volume  CommandVolumeConfig  `json:"volume"`
	Network CommandNetworkConfig `json:"network"`
	Meme    CommandMemeConfig    `json:"meme"`
	Wallet  CommandWalletConfig  `json:"wallet"`
	Help    CommandHelpConfig    `json:"help"`
}

// CommandHelpConfig holds the information about the help config
type CommandHelpConfig struct {
	Base CommandBaseConfig `json:"base"`
}

// CommandPriceConfig holds the information about the price config
type CommandPriceConfig struct {
	Base            CommandBaseConfig `json:"base"`
	TradeogreTicker string            `json:"tradeogreTicker"`
	ShowBTC         bool              `json:"showBtc"`
	ShowUSD         bool              `json:"showUsd"`
	ShowUGX         bool              `json:"showUgx"`
}

// CommandNetworkConfig holds the information about the network config
type CommandNetworkConfig struct {
	Base CommandBaseConfig `json:"base"`
	URL  string            `json:"url"`
}

// CommandVolumeConfig holds the information about the volume config
type CommandVolumeConfig struct {
	Base CommandBaseConfig `json:"base"`
}

// CommandMemeConfig holds the information about the meme config
type CommandMemeConfig struct {
	Base CommandBaseConfig `json:"base"`
}

// CommandWalletConfig holds the information about the wallet config
type CommandWalletConfig struct {
	Base CommandBaseConfig `json:"base"`
}

// CommandBaseConfig holds the common information about a command
type CommandBaseConfig struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}
