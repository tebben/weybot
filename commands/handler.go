package commands

import (
	"github.com/bwmarrin/discordgo"
)

var (
	// Commands contains the available commands
	Commands map[string]ICommand
)

// Init initialises the command, call after reading config
func Init() {
	Commands = make(map[string]ICommand, 0)

	help := NewHelpCommand()
	Commands[help.GetListener()] = help

	price := NewPriceCommand()
	Commands[price.GetListener()] = price

	volume := NewVolumeCommand()
	Commands[volume.GetListener()] = volume

	network := NewNetworkCommand()
	Commands[network.GetListener()] = network

	meme := NewMemeCommand()
	Commands[meme.GetListener()] = meme

	wallet := NewWalletCommand()
	Commands[wallet.GetListener()] = wallet
}

// HandleCommand sets the incoming command to the right ICommand
func HandleCommand(command string, params []string, s *discordgo.Session, m *discordgo.MessageCreate) {
	var c ICommand
	var found bool

	if c, found = Commands[command]; !found {
		return
	}

	c.Handle(s, m, params)
}
